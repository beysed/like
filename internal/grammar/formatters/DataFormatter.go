package formatters

type DataFormatter interface {
	Format(input any) (string, error)
}
