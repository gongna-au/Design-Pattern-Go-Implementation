package facade

import "fmt"

// 外观模式
// 无需让代码直接与数十个框架类一起工作，
// 而是创建一个封装该功能并将其隐藏在其余代码中的外观类。
// 这种结构还可以帮助最大程度地减少升级到框架的未来版本或将其替换为另一个版本的工作量。您需要在应用程序中更改的唯一一件事就是外观方法的实现。
type VideoConverter struct {
}

func NewVideoConverter() *VideoConverter {
	return &VideoConverter{}
}

func (v *VideoConverter) ConvertVideo(filename string, formate string) string {
	file := NewVideoFile(filename)
	sourceCodec := NewCodeFactory().Extract(file)
	var destinationCodec ICompression
	if formate == "mp4" {
		destinationCodec = NewMPEG4CompressionCodec()
	} else {
		destinationCodec = NewOggCompression()
	}
	buffer := NewBitrateReader().Read(file, sourceCodec)
	result := NewBitrateReader().Convert(buffer, destinationCodec)
	return result

}

// 这些是复杂的第 3 方框架
// 转换框架的一些类。我们不控制该代码，因此无法简化这些复杂的类。

type AudioMixer struct {
}

type VideoFile struct {
}

func NewVideoFile(filename string) string {
	return "file"
}

type BitrateReader struct {
}

func NewBitrateReader() *BitrateReader {
	return &BitrateReader{}
}

func (b *BitrateReader) Read(filename string, sourceCodec string) string {
	return "buffer"
}

func (b *BitrateReader) Convert(buffer string, destinationCodec ICompression) string {
	return "result"
}

type CodeFactory struct {
}

func NewCodeFactory() *CodeFactory {
	return &CodeFactory{}
}

func (c *CodeFactory) Extract(file string) string {
	return "sourceCodec"
}

type ICompression interface {
	Save()
}

type OggCompression struct {
}

func NewOggCompression() *OggCompression {
	return &OggCompression{}
}

func (o *OggCompression) Save() {

}

type MPEG4CompressionCodec struct {
}

func NewMPEG4CompressionCodec() *MPEG4CompressionCodec {
	return &MPEG4CompressionCodec{}
}

func (m *MPEG4CompressionCodec) Save() {

}

func Client() {
	convertor := NewVideoConverter()
	result := convertor.ConvertVideo("funny-cats", "mp4")
	fmt.Println(result)
}
