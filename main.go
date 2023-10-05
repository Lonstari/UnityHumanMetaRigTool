package main

import (
	//"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	T := readFile("T.fbx.meta")
	Runing := readFile("T@Running.fbx.meta")
	t_ := unity_meta{}
	err := yaml.Unmarshal(T, &t_)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	r_ := unity_meta{}
	err = yaml.Unmarshal(Runing, &r_)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//t_ r_
	kko := ""
	for i, s := range r_.ModelImporter.HumanDescription.Skeleton {
		_skeleton := findSkeleton(t_, s.Name)
		if _skeleton != nil {
			r_.ModelImporter.HumanDescription.Skeleton[i].Position = _skeleton.Position
			r_.ModelImporter.HumanDescription.Skeleton[i].Rotation = _skeleton.Rotation
			r_.ModelImporter.HumanDescription.Skeleton[i].Scale = _skeleton.Scale
			kko += SetData(s, *_skeleton) + "\n"
			fmt.Println(_skeleton.Name, " Position: ", _skeleton.Position, " Rotation: ", _skeleton.Rotation, " Scale: ", _skeleton.Scale)
		}
	}
	//write file
	err = os.WriteFile("T@Running.fbx.txt", []byte(kko), 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println("completed!")

}

func SetData(mskn Skeleton, skn Skeleton) string {
	px := strconv.FormatFloat(skn.Position.X, 'f', -1, 64)
	py := strconv.FormatFloat(skn.Position.Y, 'f', -1, 64)
	pz := strconv.FormatFloat(skn.Position.Z, 'f', -1, 64)

	rx := strconv.FormatFloat(skn.Rotation.X, 'f', -1, 64)
	ry := strconv.FormatFloat(skn.Rotation.Y, 'f', -1, 64)
	rz := strconv.FormatFloat(skn.Rotation.Z, 'f', -1, 64)
	rw := strconv.FormatFloat(skn.Rotation.W, 'f', -1, 64)

	return `    - name: ` + mskn.Name + `
      parentName: ` + mskn.ParentName.(string) + `
      position: {x: ` + px + `, y: ` + py + `, z: ` + pz + `}
      rotation: {x: ` + rx + `, y: ` + ry + `, z: ` + rz + `, w: ` + rw + `}
      scale: {x: 1, y: 1, z: 1}`
}

func findSkeleton(target_meta unity_meta, name string) *Skeleton {
	for _, s := range target_meta.ModelImporter.HumanDescription.Skeleton {
		if s.Name == name {
			return &s
		}
	}
	return nil
}

func readFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, _ := io.ReadAll(file)
	return content
}
