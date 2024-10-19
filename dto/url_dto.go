package dto

type CreateUrlRequestDto struct {
	Url string `json:"url"`
}

type CreateUrlResponseDto struct {
	ShortUrl string `json:"shortUrl"`
}
type UnwrapUrlResponseDto struct {
	OriginalUlr string `json:"originalUrl"`
}
