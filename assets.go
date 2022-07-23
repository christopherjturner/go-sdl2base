package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	_ "image/png"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func LoadTextures(textureDir string, textures map[string]*sdl.Texture) error {
	dir, err := ioutil.ReadDir(textureDir)
	if err != nil {
		return err
	}

	for _, file := range dir {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".png") {

			textureName := strings.TrimSuffix(file.Name(), ".png")
			fmt.Printf("Loading texture: %s\n", textureName)

			pic, err := img.LoadTexture(render, "./images/"+file.Name())
			if err != nil {
				return err
			}
			textures[textureName] = pic
		}
	}

	return nil
}

func UnloadTextures(textures map[string]*sdl.Texture) {
	for filename := range textures {
		textures[filename].Destroy()
		delete(textures, filename)
	}
}
