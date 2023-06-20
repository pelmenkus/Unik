#include <iostream>
#include <fstream>
#include <filesystem>
#include <string>
#include <regex>

std::string removeHTMLTags(const std::string& text) {
    // Удаление HTML-тегов с использованием регулярных выражений
    std::regex tagsRegex("<[^>]+>");
    std::string result = std::regex_replace(text, tagsRegex, "");

    // Замена обозначений &lt; и &gt; на < и >
    std::regex ltRegex("&lt;");
    std::regex gtRegex("&gt;");
    result = std::regex_replace(result, ltRegex, "<");
    result = std::regex_replace(result, gtRegex, ">");

    return result;
}

void processHTMLFiles(const std::string& directoryPath) {
    std::filesystem::path dir(directoryPath);

    // Проверка существования указанного каталога
    if (!std::filesystem::exists(dir)) {
        std::cout << "net kataloga takogo" << std::endl;
        return;
    }

    // Проверка, является ли указанный путь директорией
    if (!std::filesystem::is_directory(dir)) {
        std::cout << "ne katalog v puti" << std::endl;
        return;
    }

    // Итерация по файлам в указанной директории
    for (const auto& entry : std::filesystem::directory_iterator(dir)) {
        if (entry.is_regular_file()) {
            std::filesystem::path filePath = entry.path();
            std::string fileExtension = filePath.extension().string();

            // Проверка расширения файла
            if (fileExtension == ".html") {
                // Открытие и чтение HTML-файла
                std::ifstream inputFile(filePath);
                if (inputFile.is_open()) {
                    std::string fileContents((std::istreambuf_iterator<char>(inputFile)),
                        std::istreambuf_iterator<char>());

                    // Очистка текста от HTML-тегов
                    std::string cleanedText = removeHTMLTags(fileContents);

                    // Формирование имени для сохранения очищенного текста
                    std::string outputFileName = filePath.stem().string() + ".txt";
                    std::filesystem::path outputPath = dir / outputFileName;

                    // Сохранение очищенного текста в файле с расширением .txt
                    std::ofstream outputFile(outputPath);
                    if (outputFile.is_open()) {
                        outputFile << cleanedText;
                        outputFile.close();
                        std::cout << "file \"" << outputPath.filename() << "\" created." << std::endl;
                    }
                    else {
                        std::cout << "i cant create file \"" << outputPath.filename() << "\"." << std::endl;
                    }
                }
                else {
                    std::cout << "i cant open file \"" << filePath.filename() << "\"." << std::endl;
                }
            }
        }
    }
}

using namespace std;

int main(int argc, char* argv[]) {
    if (argc < 2) {
        string directory;
        cout << "show me the way"<<'\n';
        cin >> directory;
        processHTMLFiles(directory);
        return 0;
    }
}