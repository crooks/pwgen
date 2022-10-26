package config

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

type Flags struct {
	Config string
}

type Config struct {
	Standard struct {
		MaxSepLength     int    `yaml:"max_sep_length"`
		MaxWordLength    int    `yaml:"max_word_length"`
		MinSepLength     int    `yaml:"min_sep_length"`
		MinWordLength    int    `yaml:"min_word_length"`
		NumPasswords     int    `yaml:"num_passwords"`
		SuffixSepLength  int    `yaml:"suffix_sep_length"`
		SymbolsList      string `yaml:"symbols"`
		WordsFile        string `yaml:"words_file"`
		WordsPerPassword int    `yaml:"words_per_password"`
	} `yaml:"standard"`
	Short struct {
		MaxSepLength     int    `yaml:"max_sep_length"`
		MaxWordLength    int    `yaml:"max_word_length"`
		MinSepLength     int    `yaml:"min_sep_length"`
		MinWordLength    int    `yaml:"min_word_length"`
		NumPasswords     int    `yaml:"num_passwords"`
		SuffixSepLength  int    `yaml:"suffix_sep_length"`
		SymbolsList      string `yaml:"symbols"`
		WordsFile        string `yaml:"words_file"`
		WordsPerPassword int    `yaml:"words_per_password"`
	} `yaml:"short"`
}

// ParseConfig imports a yaml formatted config file into a Config struct
func ParseConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}

// WriteConfig will create a YAML formatted config file from a Config struct
func (c *Config) WriteConfig(filename string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// parseFlags processes arguments passed on the command line in the format
// standard format: --foo=bar
func ParseFlags() *Flags {
	f := new(Flags)
	flag.StringVar(&f.Config, "config", "njmon_exporter.yml", "Path to njmon_exporter configuration file")
	flag.Parse()
	return f
}
