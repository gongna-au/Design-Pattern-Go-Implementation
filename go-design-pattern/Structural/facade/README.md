### 外观模式的结构：

#### 1.VideoConverter 的 ConvertVideo()封装了具体的转化细节

```go
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

```

#### 2.复杂的第三方库

```go
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

```

#### 3.客户端只需要调用 VideoConverter 就可以实现转化

```go
func Client() {
	convertor := NewVideoConverter()
	result := convertor.ConvertVideo("funny-cats", "mp4")
	fmt.Println(result)
}
```
