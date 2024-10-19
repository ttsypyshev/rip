package backend

import (
	"strings"
	"time"
)

/// (Lang)
// Files
// 	ID_file 1
// 		Name
// 		Description
// 	ID_file 2
// 		Name
// 		Description
// 	ID_file 3
// 		Name
// 		Description

/// (Project)
// Projects
// 	ID_project 1
// 		Time

// 	ID_project 2
// 		Time

// MM
// 	ID_file
// 	ID_project
// 	Filename
// 	Code

// Lang представляет язык программирования.
type Lang struct {
	ID               int       `gorm:"primaryKey"`
	Name             string    `gorm:"size:255"`
	ImgLink          string    `gorm:"size:255"`
	ShortDescription string    `gorm:"size:255"`
	Author           string    `gorm:"size:255"`
	Year             string    `gorm:"size:4"`
	Version          string    `gorm:"size:50"`
	Description      string    `gorm:"type:text"`
	List             string    `gorm:"type:text"`
}

// Project представляет проект.
type Project struct {
	ID           int       `gorm:"primaryKey"`
	CreationDate time.Time
	CreationTime time.Time
	Status       int
}

// File представляет файл проекта.
type File struct {
	ID         int       `gorm:"primaryKey"`
	IDLang     int       `gorm:"column:id_lang"`
	IDProject  int       `gorm:"column:id_project"`
	Code       string    `gorm:"type:text"`
}


var Langs = []Lang{
	{
		ID:               0,
		Name:             "Python",
		ImgLink:          "http://localhost:9000/code-inspector/python.png",
		ShortDescription: "Объединяет простоту и мощь",
		Author:           "Гвидо ван Россум",
		Year:             "1991",
		Version:          "Python 3.12.6 (Sep 12, 2024)",
		Description:      "— это высокоуровневый язык программирования общего назначения, который широко используется благодаря своей гибкости, простоте и мощным возможностям расширения. Вот основные технические характеристики Python:",
		List: `Исполняемость: Python интерпретируемый язык. Код выполняется интерпретатором, который читает и исполняет команды строка за строкой.
		Модульность и библиотеки: Python поддерживает создание модулей и пакетов, а также имеет обширную стандартную библиотеку с множеством функциональных возможностей.
		Мультисистемность: Python доступен для большинства операционных систем, включая Windows, macOS и Linux.
		Объектно-Ориентированное Программирование (ООП): Python поддерживает ООП, позволяя создавать классы и объекты.
		Пакетный менеджер: Python использует pip в качестве стандартного инструмента для установки и управления пакетами и библиотеками.
		Производительность: Python может быть медленнее по сравнению с компилируемыми языками, но его производительность можно улучшить с помощью различных оптимизаций, таких как Cython и PyPy.
		Синтаксис: Python имеет простой и читаемый синтаксис, что делает его удобным для новичков. Отступы используются для обозначения блоков кода.
		Типизация: Python является динамически типизированным языком, что означает, что типы переменных проверяются во время выполнения программы.`,
	},
	{
		ID:               1,
		Name:             "C++",
		ImgLink:          "http://localhost:9000/code-inspector/cpp.png",
		ShortDescription: "Контроль и производительность в одном лице",
		Author:           "Бьёрн Страуструп",
		Year:             "1985",
		Version:          "C++20 (Dec 2020)",
		Description:      "— это мощный и высокопроизводительный язык программирования, известный своей способностью обеспечивать низкоуровневый доступ к памяти и поддерживать сложные структуры данных. Вот основные технические характеристики C++:",
		List: `Производительность: C++ обеспечивает высокую производительность за счет компиляции в машинный код и эффективного управления ресурсами.
		Синтаксис: C++ поддерживает широкий спектр программных конструкций, включая функции, классы, шаблоны и перегрузку операторов, что делает его подходящим для различных типов программирования.
		Объектно-Ориентированное Программирование (ООП): C++ позволяет создавать и использовать классы и объекты, поддерживает наследование, полиморфизм и инкапсуляцию.
		Шаблоны: C++ поддерживает шаблоны для создания обобщенного кода, который может работать с различными типами данных.
		Мультисистемность: C++ доступен для множества операционных систем и платформ, что позволяет создавать кроссплатформенные приложения.
		Стандартная библиотека: C++ включает стандартную библиотеку, которая предоставляет функции и классы для работы с коллекциями, вводом/выводом, и алгоритмами.
		Управление памятью: C++ позволяет явное управление памятью с помощью указателей и динамического выделения, что требует от программиста внимательного контроля ресурсов.
		Проектирование и сборка: C++ проектирование и сборка поддерживаются различными системами сборки и управления зависимостями, такими как CMake и Make.`,
	},
	{
		ID:               2,
		Name:             "GO",
		ImgLink:          "http://localhost:9000/code-inspector/golang.png",
		ShortDescription: "Эффективный для масштабируемых решений",
		Author:           "Роберт Гризмер, Роб Пайк",
		Year:             "2009",
		Version:          "Go 1.21.0 (Aug 2023)",
		Description:      "— это язык программирования, разработанный Google с акцентом на простоту, производительность и эффективное параллельное выполнение. Вот основные технические характеристики Go:",
		List: `Типизация: Go является статически типизированным языком, что означает проверку типов переменных на этапе компиляции.
		Исполняемость: Go компилируемый язык. Код Go компилируется в машинный код, что обеспечивает высокую производительность выполнения программ.
		Синтаксис: Go имеет простой и лаконичный синтаксис, который делает его легким для изучения и использования. Отступы используются для форматирования кода, что способствует чистоте и читаемости кода.
		Параллелизм: Go предоставляет встроенную поддержку для параллельного выполнения через горутины и каналы, что упрощает разработку многопоточных приложений.
		Мультисистемность: Go поддерживает кроссплатформенность и доступен для большинства операционных систем, включая Windows, macOS и Linux.
		Модульность и библиотеки: Go имеет встроенную поддержку для работы с пакетами и модулями, а также поставляется с обширной стандартной библиотекой.
		Пакетный менеджер: Go использует собственный инструмент go для управления пакетами и зависимостями, который упрощает процесс сборки и установки.
		Производительность: Go обеспечивает высокую производительность благодаря компиляции в машинный код и оптимизированному управлению памятью, что делает его эффективным для сетевых и серверных приложений.`,
	},
	{
		ID:               3,
		Name:             "HTML",
		ImgLink:          "http://localhost:9000/code-inspector/html.png",
		ShortDescription: "Основа структуры и содержания веб-страниц",
		Author:           "Тим Бернерс-Ли",
		Year:             "1993",
		Version:          "HTML5 (Oct 2014)",
		Description:      "— это стандартный язык разметки, используемый для создания и структурирования веб-страниц. Вот основные технические характеристики HTML:",
		List: `Разметка: HTML использует теги для определения структуры и содержания веб-страниц. Теги обозначают различные элементы, такие как заголовки, параграфы, ссылки и изображения.
		Исполняемость: HTML не является исполняемым языком программирования; вместо этого он интерпретируется веб-браузерами, которые отображают содержимое страницы согласно разметке.
		Синтаксис: HTML имеет простой и гибкий синтаксис, который легко читается и пишется. Теги обычно имеют открывающую и закрывающую форму, например <p> и </p> для параграфов.
		Структура: HTML обеспечивает основу для веб-страниц, включая такие элементы, как заголовок, тело и разделение на секции. Он также поддерживает вложенность элементов для создания сложных структур.
		Мультисистемность: HTML является стандартом для веб-разработки и поддерживается всеми современными веб-браузерами, независимо от операционной системы.
		Модульность и расширения: HTML можно расширять с помощью CSS (Cascading Style Sheets) для стилизации и JavaScript для интерактивности и динамического поведения.
		Пакетный менеджер: HTML сам по себе не имеет пакетного менеджера, но может быть использован вместе с инструментами и фреймворками для упрощения веб-разработки, такими как npm (Node Package Manager) для JavaScript-библиотек.
		Производительность: Поскольку HTML представляет собой статический язык разметки, его производительность в основном зависит от браузера и качества кода, используемого для поддержки стилей и сценариев.`,
	},
	{
		ID:               4,
		Name:             "CSS",
		ImgLink:          "http://localhost:9000/code-inspector/css.png",
		ShortDescription: "Создает оформление веб-интерфейсов",
		Author:           "Грабриел Маззоне",
		Year:             "1996",
		Version:          "CSS3 (2011)",
		Description:      "— это язык стилей, используемый для управления внешним видом и форматированием веб-страниц, созданных на HTML. Вот основные технические характеристики CSS:",
		List: `Стилевое оформление: CSS применяется для стилизации HTML-элементов, позволяя задавать цвета, шрифты, отступы, размеры и другие визуальные свойства.
		Синтаксис: CSS имеет простой синтаксис, состоящий из селекторов и деклараций. Селекторы указывают элементы, к которым применяются стили, а декларации определяют свойства и их значения.
		Каскадность и наследование: CSS поддерживает каскадность, что означает, что стили могут наследоваться от родительских элементов и могут переопределяться более специфичными правилами.
		Мультисистемность: CSS поддерживается всеми современными веб-браузерами, что позволяет создавать кроссбраузерные и адаптивные веб-дизайны.
		Модульность и расширения: CSS можно разделить на отдельные файлы и модули для упрощения управления стилями. Также существуют препроцессоры, такие как Sass и LESS, которые добавляют дополнительные возможности и упрощают разработку.
		Пакетный менеджер: CSS не имеет стандартного пакетного менеджера, но может быть интегрирован с инструментами сборки и управления зависимостями, такими как npm и Yarn.
		Производительность: Производительность CSS зависит от сложности стилей и способа их применения. Правильное использование каскадности и оптимизация кода помогают улучшить время загрузки и рендеринга веб-страниц.`,
	},
}

var Projects = []Project{
	{
		ID:           0,
		CreationTime: time.Date(2024, 9, 5, 2, 0, 0, 0, time.UTC),
		Status:       2,
	},
	{
		ID:           1,
		CreationTime: time.Date(2024, 9, 5, 2, 0, 0, 0, time.UTC),
		Status:       0,
	},
}

var Files = []File{
	{
		ID:         0,
		IDLang:    1,
		IDProject: 0,
		Code: `
#include <iostream>

int main() {
	// Объявляем переменные для хранения чисел
	int num1, num2, sum;

	// Запрашиваем у пользователя первое число
	std::cout << "Введите первое число: ";
	std::cin >> num1;

	// Запрашиваем у пользователя второе число
	std::cout << "Введите второе число: ";
	std::cin >> num2;

	// Вычисляем сумму
	sum = num1 + num2;

	// Выводим результат
	std::cout << "Сумма " << num1 << " и " << num2 << " равна " << sum << std::endl;

	return 0;
}
`,
	},
	{
		ID:         1,
		IDLang:    0,
		IDProject: 1,
		Code: `
from flask import Flask, render_template

app = Flask(__name__)

@app.route('/')
def home():
	return render_template('index.html')

if __name__ == '__main__':
	app.run(debug=True)
`,
	},
	{
		ID:         2,
		IDLang:    3,
		IDProject: 1,
		Code: `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel="stylesheet" href="{{ url_for('static', filename='index.css') }}">
	<title>Простой Flask проект</title>
</head>
<body>
	<div class="container">
		<h1>Добро пожаловать в мой проект!</h1>
		<p>Это простой проект на Flask.</p>
	</div>
</body>
</html>
`,
	},
	{
		ID:         3,
		IDLang:    4,
		IDProject: 1,
		Code: `
body {
	font-family: Arial, sans-serif;
	background-color: #f0f0f0;
	margin: 0;
	padding: 0;
}

.container {
	max-width: 600px;
	margin: 50px auto;
	padding: 20px;
	background-color: white;
	box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
	border-radius: 8px;
	text-align: center;
}

h1 {
	color: #333;
}

p {
	color: #666;
}
`,
	},
}

// Общие функции получения данных
func getByID[T any](items []T, id int, idSelector func(T) int) (T, bool) {
	for _, item := range items {
		if idSelector(item) == id {
			return item, true
		}
	}
	var empty T
	return empty, false
}

func getAll[T any](items []T) []T {
	return items
}

// Получение всех сущностей
func GetLangs() []Lang {
	return getAll(Langs)
}

func GetProjects() []Project {
	return getAll(Projects)
}

func GetFiles() []File {
	return getAll(Files)
}

// Получение сущностей по ID
func GetLangByID(langID int) (Lang, bool) {
	return getByID(Langs, langID, func(lang Lang) int { return lang.ID })
}

func GetProjectByID(projectID int) (Project, bool) {
	return getByID(Projects, projectID, func(project Project) int { return project.ID })
}

func GetFileByID(fileID int) (File, bool) {
	return getByID(Files, fileID, func(file File) int { return file.ID })
}

// Получение файлов для конкретного проекта
func GetFilesForProject(projectID int) []File {
	var matchedFiles []File
	for _, file := range Files {
		if file.IDProject == projectID {
			matchedFiles = append(matchedFiles, file)
		}
	}
	return matchedFiles
}

// Фильтрация языков по запросу
func FilterLangsByQuery(langs []Lang, query string) []Lang {
	var filteredLangs []Lang
	lowerQuery := strings.ToLower(query)

	for _, lang := range langs {
		if strings.Contains(strings.ToLower(lang.Name), lowerQuery) {
			filteredLangs = append(filteredLangs, lang)
		}
	}
	return filteredLangs
}

// func FindMaxProjectID() int {
// 	maxID := -1
// 	for _, project := range Projects {
// 		if project.ID > maxID {
// 			maxID = project.ID
// 		}
// 	}
// 	return maxID
// }

// func GetLangsForProject(matchedFiles []File, projectID int) []Lang {
// 	var matchedLangs []Lang
// 	for _, file := range matchedFiles {
// 		log.Println(file)
// 		if file.ID_lang < 0 || file.ID_lang >= len(Langs) {
// 			log.Printf("Invalid ID_lang: %d for file %v", file.ID_lang, file)
// 			continue
// 		}
// 		matchedLangs = append(matchedLangs, Langs[file.ID_lang])
// 	}
// 	return matchedLangs
// }
