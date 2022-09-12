package main
import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"bufio"

	"github.com/kkdai/youtube/v2"  //package youtube banyak lah
	
)

// URL REGEX untuk check validasi url youtubenya
var (
	urlregex = regexp.MustCompile(`(http:|https:)?\/\/(www\.)?(youtube.com|youtu.be)\/(watch)?(\?v=)?(\S+)?`)
)

//Fungsi download dari package tadi
func Download(url string) {
	if !urlregex.MatchString(url) {
		fmt.Println("URL tidak valid")
	}
	
	videoUrls := url
	client := youtube.Client{}

	//Check error
	video, err := client.GetVideo(videoUrls)
	if err != nil {
		log.Println(err)
	}

	formats := video.Formats.WithAudioChannels()
	stream, _, err := client.GetStream(video, &formats[2])
	if err != nil {
		log.Println(err)
	}

	//konversi ke mp3
	hasilDownload := strings.ReplaceAll(video.Title+".mp3", "/", "|")
	fmt.Println("Judul :", hasilDownload)
	fmt.Println("File :", video.Title+ " Sedang di download, harap tunggu")
	file, err := os.Create(hasilDownload)
	
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	_, err = io.Copy(file, stream)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Selesai di download ", hasilDownload)
}

func main() {
	//Membaca link
    reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukan url: ")
	
	urlnew, _ := reader.ReadString('\n')
	
	//Eksekusi fungsi download
	Download(urlnew)

}
