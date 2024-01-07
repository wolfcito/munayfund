package ipfs

import (
	"fmt"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

type IPFSService struct {
	shell *shell.Shell
}

func NewIPFSService() *IPFSService {
	return &IPFSService{
		shell: shell.NewShell(os.Getenv("IPFSURL")),
	}
}

func (s *IPFSService) UploadFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Subir el archivo al nodo de IPFS
	cid, err := s.shell.Add(file)
	if err != nil {
		return "", err
	}

	// Construir la URL del archivo utilizando el CID
	url := fmt.Sprintf("https://ipfs.io/ipfs/%s", cid)

	return url, nil
}
