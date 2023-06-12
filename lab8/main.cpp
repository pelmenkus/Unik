#include <iostream>
#include "equality.h"
using namespace std;
int main() {
    const int L = 0;  // Нижняя граница диапазона коэффициентов
    const int H = 10; // Верхняя граница диапазона коэффициентов

    const int N = 3;  // Количество переменных (размер вектора)

    Equality<L, N> equation;
    equation.setCoefficient(0, 2);   // a1 = 2
    equation.setCoefficient(1, 3);   // a2 = 3
    equation.setCoefficient(2, 1);   // a3 = 1
    equation.setConstant(10);       // b = 10

    int values[N] = {1, 2, 3};  // Значения переменных x1, x2, x3

    cout<<"Constant = 10"<<'\n';

    bool result = equation.checkEquality(values);
    if (result) {
        std::cout << "Correct" << std::endl;
    } else {
        std::cout << "not Correct" << std::endl;
    }

    cout<<"Constant = 11"<<'\n';

    equation.setConstant(11);
    result = equation.checkEquality(values);
    if (result) {
        std::cout << "Correct" << std::endl;
    } else {
        std::cout << "not Correct" << std::endl;
    }

    return 0;
}
