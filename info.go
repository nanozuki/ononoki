package akko

import "github.com/getkin/kin-openapi/openapi3"

type InfoBuilder struct {
	info *openapi3.Info
}

func Info() *InfoBuilder {
	return &InfoBuilder{}
}

func (b *InfoBuilder) Description(desc string) *InfoBuilder {
	b.info.Description = desc
	return b
}

func (b *InfoBuilder) TermOfService(tos string) *InfoBuilder {
	b.info.TermsOfService = tos
	return b
}

func (b *InfoBuilder) Contact(contact *ContactBuilder) *InfoBuilder {
	b.info.Contact = contact.contact
	return b
}

func (b *InfoBuilder) License(license *LicenseBuilder) *InfoBuilder {
	b.info.License = license.license
	return b
}

func (b *InfoBuilder) AttachToAPI(api *OpenAPIBuilder) *InfoBuilder {
	b.info.Title = api.api.Info.Title
	b.info.Version = api.api.Info.Version

	api.api.Info = b.info
	return b
}
