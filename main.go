package main

/* rpgsave-decrypt
cli tool to decompress RPG Maker MV save files.

`.rpgsave` files are compressed using lz-string, then encoded in base64.
@daku10's go-lz-string library is used to decompress the save file.

usage:
	either:
		1. drag and drop the save file onto the executable
		2. rpgsave-decode.exe <save file path>
*/

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/log"
	lzstring "github.com/daku10/go-lz-string" // migrated from @pieroxy/lz-string-go upon issues with decoding (issues with int -> string conversion?)
)

func main() {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
	})
	// set log level
	logger.SetLevel(log.DebugLevel)
	logger.Info("starting rpgsave-decode")
	// get file path from command line args
	args := os.Args[1:]
	if len(args) == 0 {
		logger.Error("no file path provided")
		logger.Info("press enter to exit")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		os.Exit(1)
	}
	filePath := args[0]
	logger.Print("file path: " + filePath)

	// check if file is ".rpgsave"
	if filepath.Ext(filePath) == ".rpgsave" {
		decodeRpgSave(filePath, logger)
		return
	}
	if filepath.Ext(filePath) == ".json" {
		encodeRpgSave(filePath, logger)
		return
	}
	logger.Error("invalid file type")
	return
}

func decodeRpgSave(filePath string, logger *log.Logger) {
	// open file
	file, err := os.Open(filePath)
	if err != nil {
		logger.Error("error opening file")
		logger.Error(err)
		logger.Info("press enter to exit")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		os.Exit(1)
	}
	// read file
	fileInfo, err := file.Stat()
	if err != nil {
		logger.Error("error reading file")
		logger.Error(err)
		logger.Info("press enter to exit")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		os.Exit(1)
	}
	fileSize := fileInfo.Size()
	logger.Print("file size: " + fmt.Sprint(fileSize))
	fileBytes := make([]byte, fileSize)
	_, err = file.Read(fileBytes)
	if err != nil {
		logger.Error("error reading file")
		logger.Error(err)
		logger.Info("press enter to exit")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		os.Exit(1)
	}

	// decompress
	logger.Info("decompressing")
	decompressedBytes, err := lzstring.DecompressFromBase64(string(fileBytes))
	// write to file
	logger.Info("writing to file")
	fileName := filepath.Base(filePath)
	fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))
	fileName += ".json"
	logger.Print("file name: " + fileName)

	if _, err := os.Stat(fileName); err == nil {
		logger.Warn("file", fileName, " exits, overwrite? (y/n)")
		overwrite, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		overwrite = strings.TrimSpace(overwrite)
		if overwrite == "y" {
			logger.Info("trunicating file", fileName)
			err := os.Remove(fileName)
			if err != nil {
				logger.Error("error deleting file")
				logger.Error(err)
				logger.Info("press enter to exit")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				os.Exit(1)
			}

		} else {
			logger.Info("not overwriting")
			logger.Info("press enter to exit")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			os.Exit(0)
		}
	}

	// pretty print json
	// logger.Info("pretty printing json")
	// var jsonData []byte
	// // json Indent
	// jsonData, err = json.MarshalIndent(decompressedBytes, "", "  ")
	// if err != nil {
	// 	logger.Error("error marshaling JSON")
	// 	logger.Error(err)
	// 	logger.Info("press enter to exit")
	// 	bufio.NewReader(os.Stdin).ReadBytes('\n')
	// 	os.Exit(1)
	// }
	// // unquote
	// jsonDataStr, err := strconv.Unquote(string(jsonData))
	// if err != nil {
	// 	logger.Error("error unquoting JSON")
	// 	logger.Error(err)
	// 	logger.Info("press enter to exit")
	// 	bufio.NewReader(os.Stdin).ReadBytes('\n')
	// 	os.Exit(1)
	// }
	// jsonData = []byte(jsonDataStr)

	// // Indent
	// var indentedData bytes.Buffer
	// err = json.Indent(&indentedData, jsonData, "", "  ")
	// if err != nil {
	// 	logger.Error("error marshaling JSON")
	// 	logger.Error(err)
	// 	logger.Info("press enter to exit")
	// 	bufio.NewReader(os.Stdin).ReadBytes('\n')
	// 	os.Exit(1)
	// }
	// jsonData = indentedData.Bytes()

	//str to bytes
	jsonData := []byte(decompressedBytes)

	// write to file
	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		logger.Error("error writing to file")
		logger.Error(err)
		logger.Info("press enter to exit")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		os.Exit(1)
	}
	logger.Info("done!")
	logger.Info("press enter to exit")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func encodeRpgSave(filePath string, logger *log.Logger) {
	// open file
	file, err := os.Open(filePath)
	if err != nil {
		logger.Error("error opening file")
		logger.Error(err)
		logger.Info("press enter to exit")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		os.Exit(1)
	}
	// read file
	fileInfo, err := file.Stat()
	if err != nil {
		logger.Error("error reading file")
		logger.Error(err)
		logger.Info("press enter to exit")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		os.Exit(1)
	}
	fileSize := fileInfo.Size()
	logger.Print("file size: " + fmt.Sprint(fileSize))
	fileBytes := make([]byte, fileSize)
	_, err = file.Read(fileBytes)
	if err != nil {
		logger.Error("error reading file")
		logger.Error(err)
		logger.Info("press enter to exit")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		os.Exit(1)
	}

	// unpretty print json
	logger.Info("unpretty printing json")
	// quote
	// jsonDataStr, err := strconv.QuotedPrefix(string(fileBytes))
	// if err != nil {
	// 	logger.Error("error quoting JSON")
	// 	logger.Error(err)
	// 	logger.Info("press enter to exit")
	// 	bufio.NewReader(os.Stdin).ReadBytes('\n')
	// 	os.Exit(1)
	// }
	// fileBytes = []byte(jsonDataStr)

	// unindent
	// var unindentedData bytes.Buffer
	// err = json.Compact(&unindentedData, fileBytes)
	// if err != nil {
	// 	logger.Error("error unindenting JSON")
	// 	logger.Error(err)
	// 	logger.Info("press enter to exit")
	// 	bufio.NewReader(os.Stdin).ReadBytes('\n')
	// 	os.Exit(1)
	// }
	// fileBytes = unindentedData.Bytes()

	// compress
	logger.Info("compressing")
	compressedBytes, _ := lzstring.CompressToBase64(string(fileBytes))
	// write to file
	logger.Info("writing to file")
	fileName := filepath.Base(filePath)
	fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))
	fileName += ".rpgsave"
	logger.Print("file name: " + fileName)

	if _, err := os.Stat(fileName); err == nil {
		logger.Warn("file", fileName, " exits, overwrite? (y/n)")
		overwrite, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		overwrite = strings.TrimSpace(overwrite)
		if overwrite == "y" {
			logger.Info("trunicating file", fileName)
			err := os.Remove(fileName)
			if err != nil {
				logger.Error("error deleting file")
				logger.Error(err)
				logger.Info("press enter to exit")
				bufio.NewReader(os.Stdin).ReadBytes('\n')
				os.Exit(1)
			}

		} else {
			logger.Info("not overwriting")
			logger.Info("press enter to exit")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			os.Exit(0)
		}
	}

	// write to file
	err = os.WriteFile(fileName, []byte(compressedBytes), 0644)
	if err != nil {
		logger.Error("error writing to file")
		logger.Error(err)
		return
	}
	logger.Info("done!")
}
