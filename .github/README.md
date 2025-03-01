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
    5, 
    []byte("Hello"),
)
```
### 3. Основная структура SIF
```go title="Example"
type SIF struct {
    Header  Header
    Tags    []tag.Tag // Теги (метаданные)
}
```
#### Пример создания SIF:
```go title="Example"
content := tag.New(
    tag.ContentSignature, 
    5, 
    []byte("Hello, SIF!"),
)
metaTag :=  tag.New(
    tag.ContentSignature, 
    5, 
    []byte("Hello, SIF!"),
)
sifFile, err := sif.New(content, metaTag)
```
Что будет в файле:
```bash
83 73 70 1 0 0 0 0 1 0 0 0 5 72 101 108 108 111 44 32 83 73 70 33 1 0 0 0 5 72 101 108 108 111 44 32 83 73 70 33
```