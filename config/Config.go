package config

import net_url "net/url"
import "sort"
import "strings"

type Config struct {
	base_url *net_url.URL
	preview  bool
	roots    map[string]string
}

func NewConfig(base string, preview bool, path_to_root map[string]string) *Config {

	base_url, err := net_url.Parse(base)

	if err == nil {

		var config Config

		config.base_url = base_url
		config.preview  = preview
		config.roots    = make(map[string]string)

		for path, root := range path_to_root {
			config.AddRoot(path, root)
		}

		return &config

	}

	return nil

}

func (config *Config) AddRoot(path string, root string) bool {

	if strings.HasPrefix(path, "/") && strings.HasPrefix(root, "/") {

		config.roots[path] = root

		return true

	}

	return false

}

func (config *Config) BaseURL() *net_url.URL {
	return config.base_url
}

func (config *Config) GetPaths() []string {

	result := make([]string, 0)

	for path, _ := range config.roots {
		result = append(result, path)
	}

	sort.Strings(result)

	return result

}

func (config *Config) GetRoot(http_path string) string {

	var found_root string

	for path, root := range config.roots {

		if strings.HasPrefix(http_path, path) {
			found_root = root
			break
		}

	}

	return found_root

}

func (config *Config) LivePreview() bool {
	return config.preview == true
}

func (config *Config) ResolvePath(http_path string) string {

	var found_path string
	var found_root string

	for path, root := range config.roots {

		if strings.HasPrefix(http_path, path) {
			found_path = path
			found_root = root
			break
		}

	}

	if found_path != "" {

		relative_path := strings.TrimPrefix(http_path, found_path)

		if strings.HasPrefix(relative_path, "/") {
			return found_root + relative_path
		} else {
			return found_root + "/" + relative_path
		}

	} else {
		return "/tmp" + http_path
	}

}

func (config *Config) RemoveRoot(path string) bool {

	_, ok := config.roots[path]

	if ok == true {

		delete(config.roots, path)

		return true

	}

	return false

}
