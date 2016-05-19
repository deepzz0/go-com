package useragent

type Model struct {
	Regex string
	Model string
}

type Robot struct {
	Regex    string
	Name     string
	URL      string
	Producer ProduceBy
}

type ProduceBy struct {
	Name string
	URL  string
}

type OS struct {
	Regex   string
	Name    string
	Version string
}

type Vendor struct {
	Producer string
	Regex    []string
}

///////////////////////////////////////////////////////////////////////
type Client struct {
	Regex   string
	Name    string
	Version string
}

type Browser struct {
	Client
	Engine struct {
		Default  string
		Versions map[string]string
	}
}

type Engine struct {
	Regex string
	Name  string
}

type Reader struct {
	Client
	URL  string
	Type string
}

///////////////////////////////////////////////////////////////////////
type Device struct {
	Producer string
	Regex    string
	Model    string
	Models   []Model
}
