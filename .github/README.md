# SIF-Go

![sif logo](/.github/sif.png)

## 💡 Краткое описание

`SIF-Go` — это библиотека на Go для работы с форматом SIF (Simple Information File). 

## ⚙️ Установка

Чтобы установить библиотеку, выполните следующую команду:

```bash
go get github.com/WhiCu/sif-go
```

## 📖 Структура формата SIF

Файл SIF состоит из заголовка и последовательности тегов (чанков).
Каждый тег содержит сигнатуру, длину данных и сами данные.

### 1. Заголовок (Header)
```go title="Example"
type Header struct {
    Signature [3]byte // Сигнатура формата: "SIF"
    Version   byte    // Версия формата (например, 0x01)
    Reserve   [4]byte // Зарезервированные байты
}
```
- Размер: 8 байт.

- Сигнатура: [3]byte{'S', 'I', 'F'}.

- Версия: Определяет версию спецификации.

- Резерв: Зарезервировано для будущих расширений.

### 2. Тег (Tag)
```go title="Example"
type Tag struct {
    Signature byte   // Тип тега (например, Content, Info)
    Length    int32  // Длина данных
    Data      []byte // Произвольные данные
}
```
#### Сигнатуры:

- ContentSignature: Тег с основным содержимым.

- InfoSignature: Тег с метаданными.

- Другие типы могут быть добавлены пользователем.

#### Пример тега:
```go title="Example"
// Создание тега с данными "Hello"
tag := tag.New(
    tag.ContentSignature, 
    []byte("Hello"),
)
```
### 3. Основная структура SIF
```go title="Example"
type SIF struct {
    Header  Header
    Tags    []*tag.Tag // Теги (метаданные)
}
```
#### Пример создания SIF:
```go title="Example"
content := tag.New(
    tag.ContentSignature, 
    []byte("Hello, SIF!"),
)
metaTag :=  tag.New(
    tag.ContentSignature, 
    []byte("Hello, SIF!"),
)
sifFile, err := sif.New(content, metaTag)
```
Что будет в файле:
```bash
83 73 70 1 0 0 0 0 1 0 0 0 5 72 101 108 108 111 44 32 83 73 70 33 1 0 0 0 5 72 101 108 108 111 44 32 83 73 70 33
```

## 🕶️ Предлогаемые стандартные тэги

### Тэг Info
```go title="Example"
func NewInfoTag(data []byte) *tag.Tag {
	return tag.New(
		tag.InfoSignature,
		data)
}
```

### Тэг Content
```go title="Example"
func NewContentTag(data []byte) *tag.Tag {
	return tag.New(
		tag.ContentSignature,
		data,
	)
}
```

### Тэг Number
```go title="Example"
func NewNumberTag(num int32) *tag.Tag {
	numBytes := tag.Int32ToBytes(num)
	return tag.New(
		NumSigature,
		numBytes[:],
	)
}
```

## 📝 Примеры использования

### Создание и запись SIF-файла
```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/WhiCu/sif-go/sif"
	"github.com/WhiCu/sif-go/tag/extension"
)

func main() {
	// Создание тега с содержимым
	content, err := extension.NewContentTagFromString("Hello, SIF!")
	if err != nil {
		log.Fatalf("Failed to create content tag: %v", err)
	}

	// Создание SIF-файла
	sifFile := sif.New(content)

	// Запись SIF-файла
	file, err := os.OpenFile("./ground/example.sif", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(sifFile.Bytes())
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	fmt.Println("SIF-файл успешно создан и записан.")
}
```

### Чтение SIF-файла
```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/WhiCu/sif-go/sif"
	"github.com/WhiCu/sif-go/sif/decode"
)

func main() {
	// Открытие SIF-файла
	file, err := os.Open("./ground/example.sif")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Декодирование SIF-файла
	sf := new(sif.SIF)
	err = decode.UnmarshalReader(file, sf)
	if err != nil {
		log.Fatalf("Failed to unmarshal: %v", err)
	}

	// Вывод содержимого
	fmt.Println("SIF-файл:", sf)
	fmt.Println("Содержимое:", string(sf.Tags[0].Data))
}
```

## 🏗️ Схема архитектуры

```
sif-go/
├── cmd/
│   └── main.go           # Точка входа приложения
├── sif/
│   ├── header.go         # Структура и логика заголовка SIF
│   ├── sif.go            # Основная структура SIF
│   └── decode/
│       └── decode.go     # Логика декодирования SIF
├── tag/
│   ├── tag.go            # Структура и логика тегов
│   └── extension/
│       ├── signature.go  # Константы сигнатур тегов
│       ├── tagcon.go     # Реализация тега содержимого
│       └── taginf.go     # Реализация информационного тега
└── ground/               # Папка для хранения SIF-файлов
    └── password.sif      # Пример SIF-файла
```