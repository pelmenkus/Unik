#ifndef EQUALITY_H
#define EQUALITY_H
#include<vector>;
using namespace std;
class Equality {
private:
    int N;
    double L;
    std::vector <double> coefficients;
    double constant;

public:
    Equality(double l, int n) {
        N=n;
        L=l;
        coefficients.resize(N);
        for (int i = 0; i < N; i++) {
            coefficients[i] = 0;
        }
        constant = 0;
    }

    void setCoefficient(int index, double value) {
        coefficients[index] = value - L;
    }

    void setConstant(double value) {
        constant = value - L;
        //cout<<constant;
    }

    bool checkEquality(const double* values) {
        double result = 0;
        for (int i = 0; i < N; i++) {
            result += coefficients[i] * values[i];
        }
        cout<<"result = "<<result<<'\n';
        return result == constant;
    }
};

#endif  // EQUALITY_H
