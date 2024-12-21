package example

type Example struct {
    Name  string
    Value int
}

func (e *Example) GetName() string {
    return e.Name
}

func (e *Example) GetValue() int {
    return e.Value
}

func (e *Example) SetName(name string) {
    e.Name = name
}

func (e *Example) SetValue(value int) {
    e.Value = value
}