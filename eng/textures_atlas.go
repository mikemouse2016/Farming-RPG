package eng

type texturesAtlas struct {
	handle map[string]*Texture
}

var taInstance *texturesAtlas

func TexturesAtlas() *texturesAtlas {
	if taInstance == nil {
		taInstance = &texturesAtlas{}
		taInstance.handle = make(map[string]*Texture)
	}
	return taInstance
}

func (atlas *texturesAtlas) Add(name string, filepath string, window *Window) {
	temp := new(Texture)
	if temp.LoadFromFile(filepath, window) {
		atlas.handle[name] = temp
	}
}

func (atlas *texturesAtlas) GetByName(name string) *Texture {
	return taInstance.handle[name]
}