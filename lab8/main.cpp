#include <iostream>
#include "equality.h"
using namespace std;
int main() {
    double L = 0;  // Нижняя граница диапазона коэффициентов
    double H = 10; // Верхняя граница диапазона коэффициентов

    const int N = 3;  // Количество переменных (размер вектора)

    Equality equation(L,N);
    equation.setCoefficient(0, 2.3);   // a1 = 2
    equation.setCoefficient(1, 3.1);   // a2 = 3
    equation.setCoefficient(2, 1.5);   // a3 = 1
    equation.setConstant(10);       // b = 10

    double values[N] = {1.5, 2.77, 3.8};  // Значения переменных x1, x2, x3

    cout<<"Constant = 10"<<'\n';

    bool result = equation.checkEquality(values);
    if (result) {
        std::cout << "Correct" << std::endl;
    } else {
        std::cout << "not Correct" << std::endl;
    }

    cout<<"Constant = 17.737"<<'\n';

    equation.setConstant(17.737);
    result = equation.checkEquality(values);
    if (result) {
        std::cout << "Correct" << std::endl;
    } else {
        std::cout << "not Correct" << std::endl;
    }

    return 0;
}
