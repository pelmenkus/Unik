#include <iostream>
#include <vector>

void printVector(std::vector<int>& vector)
{
    for (auto& item : vector) {
        std::cout << item << " ";
    }
    std::cout << std::endl;
    std::cout << std::endl;
}
int main(){
std::vector<int> vector(10, 14);

    printVector(vector);

    return 0;
}
