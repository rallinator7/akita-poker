//+build mage

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/sh"
)

var (
	// needed so Go properly builds binaries for alpine images
	env = map[string]string{
		"CGO_ENABLED": "0",
	}
	owner   = "rallinator7"
	baseDir = getMageDir()
)

func gitCommit() string {
	s, e := sh.Output("git", "rev-parse", "--short", "HEAD")
	if e != nil {
		fmt.Printf("Failed to get GIT version: %s\n", e)
		return ""
	}
	return s
}

func getMageDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	return dir
}

// tests the code for the server
func TestServer() error {
	os.Chdir(filepath.Join(baseDir, "server"))

	files, err := ioutil.ReadDir("./")
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	for _, f := range files {
		if f.IsDir() {
			err := sh.Run("go", "test", "./"+f.Name())
			if err != nil {
				return fmt.Errorf("server failed testing: %s", err)
			}
		}
	}

	return nil
}

// places updated protofiles for client and server
func Proto() error {
	// get current dir mage is running in
	dir := baseDir

	// point to proto file location
	protoPath := filepath.Join(dir, "proto")

	// get files in proto path
	files, err := ioutil.ReadDir(protoPath)
	if err != nil {
		return fmt.Errorf("could not get files in %s: %s", protoPath, err)
	}

	// get the generated protobuf files for each proto file for go
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".proto") {

			err = sh.Run("protoc", "--proto_path="+protoPath, "--go-grpc_out=.", file.Name())
			if err != nil {
				return fmt.Errorf("could not create go proto files: %s", err)
			}

			err = sh.Run("protoc", "--proto_path="+protoPath, "--go_out=.", file.Name())
			if err != nil {
				return fmt.Errorf("could not create go proto files: %s", err)
			}
		}
	}

	// create typescript proto
	// tsDir := filepath.Join(dir, "client")
	// err = os.Chdir(tsDir)
	// if err != nil {
	// 	return fmt.Errorf("could not change directory to client: %s", err)
	// }

	// tsProtoPath := filepath.Join(tsDir, "src", "proto")

	// err = os.Mkdir(tsProtoPath, 0755)
	// if err != nil {
	// 	fmt.Println("proto folder already exists! Continuing...")
	// }

	// gentsPath := filepath.Join(tsDir, "node_modules", ".bin", "protoc-gen-ts")

	// err = sh.Run("npx", "grpc_tools_node_protoc", "--js_out=import_style=commonjs,binary:"+tsProtoPath,
	// 	"--grpc_out="+tsProtoPath, "--plugin=protoc-gen-grpc="+gentsPath, "--proto_path="+protoPath, filepath.Join(protoPath, "hand.proto"))
	// if err != nil {
	// 	return fmt.Errorf("could not create ts proto files: %s", err)
	// }

	// err = sh.Run("npx", "grpc_tools_node_protoc", "--plugin=protoc-gen-ts="+gentsPath,
	// 	"--ts_out="+tsProtoPath, "--proto_path="+protoPath, filepath.Join(protoPath, "hand.proto"))
	// if err != nil {
	// 	return fmt.Errorf("could not create ts proto files: %s", err)
	// }

	return nil
}

type DockerImage struct {
	RelativePath string
	Name         string
	Tag          string
}

// builds the akita-poker server image
func BuildServer() error {
	// get current dir mage is running in
	dir := baseDir

	builds := []DockerImage{
		{RelativePath: "server", Name: "ghcr.io/rallinator7/akita-poker-server", Tag: gitCommit()},
	}

	for _, build := range builds {
		os.Chdir(filepath.Join(dir, build.RelativePath))

		err := sh.RunWithV(env, "go", "build", "-o", "akita-poker-server")
		if err != nil {
			return fmt.Errorf("could not build binary: %s", err)
		}
		err = sh.Run("docker", "build", "-t", fmt.Sprintf("%s:%s", build.Name, build.Tag), "-t", fmt.Sprintf("%s:%s", build.Name, "latest"), ".")
		if err != nil {
			return fmt.Errorf("could not build docker image: %s", err)
		}

		err = sh.Run("rm", "akita-poker-server")
		if err != nil {
			return fmt.Errorf("could not remove binary: %s", err)
		}
	}

	return nil
}

// builds the akita-poker client image
func BuildClient() error {
	// get current dir mage is running in
	dir := baseDir

	builds := []DockerImage{
		{RelativePath: "client", Name: "ghcr.io/rallinator7/akita-poker-client", Tag: gitCommit()},
	}

	for _, build := range builds {
		os.Chdir(filepath.Join(dir, build.RelativePath))

		err := sh.RunWithV(env, "npm", "i")
		if err != nil {
			return fmt.Errorf("could not install npm deps: %s", err)
		}
		err = sh.RunWithV(env, "npm", "run", "build")
		if err != nil {
			return fmt.Errorf("could not run npm build: %s", err)
		}
		err = sh.Run("docker", "build", "-t", fmt.Sprintf("%s:%s", build.Name, build.Tag), "-t", fmt.Sprintf("%s:%s", build.Name, "latest"), ".")
		if err != nil {
			return fmt.Errorf("could not build docker image: %s", err)
		}
	}

	return nil
}

// builds both the client and the server
func BuildAll() error {
	err := BuildServer()
	if err != nil {
		return fmt.Errorf("could not build server: %s", err)
	}

	err = BuildClient()
	if err != nil {
		return fmt.Errorf("could not build client: %s", err)
	}

	return nil
}

// pushes images to GitHub
func PushImages() error {
	images := []string{"akita-poker-client", "akita-poker-server"}
	commit := gitCommit()
	for _, image := range images {
		err := sh.Run("docker", "push", fmt.Sprintf("ghcr.io/%s/%s:%s", owner, image, commit))
		if err != nil {
			return fmt.Errorf("could not push docker image: %s", err)
		}
		err = sh.Run("docker", "push", fmt.Sprintf("ghcr.io/%s/%s:latest", owner, image))
		if err != nil {
			return fmt.Errorf("could not push docker image: %s", err)
		}
	}

	return nil
}

// ci run for GitHub Actions
func CI() error {
	err := TestServer()
	if err != nil {
		return fmt.Errorf("failed testing server: %s", err)
	}

	err = BuildServer()
	if err != nil {
		return fmt.Errorf("could not build server: %s", err)
	}

	err = BuildClient()
	if err != nil {
		return fmt.Errorf("could not build client: %s", err)
	}

	err = PushImages()
	if err != nil {
		return fmt.Errorf("could not push images: %s", err)
	}

	return nil
}
