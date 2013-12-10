package hdf5

type Location interface {
	Close() error
	Name() string
	FileName() string
	File() *File
	CreateGroup(name string) (*Group, error)
	OpenGroup(name string) (*Group, error)
	OpenDatatype(name string, tapl_id int) (*Datatype, error)
	NumObjects() (uint, error)
	ObjectNameByIndex(idx uint) (string, error)
	CreateDataset(name string, dtype *Datatype, dspace *Dataspace) (*Dataset, error)
	CreateDatasetWith(name string, dtype *Datatype, dspace *Dataspace) (*Dataset, error)
	OpenDataset(name string) (*Dataset, error)
	CreateTable(name string, dtype *Datatype, chunkSize, compression int) (*Table, error)
	CreateTableFrom(name string, dtype interface{}, chunkSize, compression int) (*Table, error)
	OpenTable(name string) (*Table, error)
}
