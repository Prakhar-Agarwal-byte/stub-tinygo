package main

import (
	"archive/tar"
	"fmt"
	"os"
)

func main() {
	// Create a tar archive file
	tarFilename := "my_archive.tar"
	tarFile, err := os.Create(tarFilename)
	if err != nil {
		fmt.Printf("Error creating tar archive: %v\n", err)
		return
	}
	defer tarFile.Close()

	// Create a new tar writer
	tarWriter := tar.NewWriter(tarFile)

	// List of files to include in the archive
	files := []struct {
		Name, Content string
	}{
		{"file1.txt", "This is the content of file1."},
		{"file2.txt", "This is the content of file2."},
	}

	// Add files to the archive
	for _, file := range files {
		header := &tar.Header{
			Name: file.Name,
			Mode: 0644, // File permissions
			Size: int64(len(file.Content)),
		}

		if err := tarWriter.WriteHeader(header); err != nil {
			fmt.Printf("Error writing tar header for %s: %v\n", file.Name, err)
			return
		}

		if _, err := tarWriter.Write([]byte(file.Content)); err != nil {
			fmt.Printf("Error writing content for %s: %v\n", file.Name, err)
			return
		}
	}

	// Close the tar writer
	if err := tarWriter.Close(); err != nil {
		fmt.Printf("Error closing tar writer: %v\n", err)
		return
	}

	fmt.Printf("Tar archive '%s' created successfully.\n", tarFilename)

	// // Extract the tar archive
	// fmt.Println("Extracting the contents of the tar archive:")
	// tarFile, err = os.Open(tarFilename)
	// if err != nil {
	// 	fmt.Printf("Error opening tar archive: %v\n", err)
	// 	return
	// }
	// defer tarFile.Close()

	// tarReader := tar.NewReader(tarFile)

	// for {
	// 	header, err := tarReader.Next()
	// 	if err == io.EOF {
	// 		break // End of archive
	// 	}
	// 	if err != nil {
	// 		fmt.Printf("Error reading tar header: %v\n", err)
	// 		return
	// 	}

	// 	fmt.Printf("Extracting %s\n", header.Name)

	// 	// Create a new file to write the extracted content
	// 	extractedFile, err := os.Create(header.Name)
	// 	if err != nil {
	// 		fmt.Printf("Error creating %s: %v\n", header.Name, err)
	// 		return
	// 	}
	// 	defer extractedFile.Close()

	// 	// Copy the content from the tar archive to the new file
	// 	if _, err := io.Copy(extractedFile, tarReader); err != nil {
	// 		fmt.Printf("Error extracting content for %s: %v\n", header.Name, err)
	// 		return
	// 	}
	// }

	// fmt.Println("Extraction completed successfully.")
}
