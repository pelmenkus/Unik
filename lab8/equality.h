#ifndef EQUALITY_H
#define EQUALITY_H
using namespace std;
template <int L, int N>
class Equality {
private:
    int coefficients[N];
    int constant;

public:
    Equality() {
        for (int i = 0; i < N; i++) {
            coefficients[i] = 0;
        }
        constant = 0;
    }

    void setCoefficient(int index, int value) {
        coefficients[index] = value - L;
    }

    void setConstant(int value) {
        constant = value - L;
    }

    bool checkEquality(const int* values) {
        int result = 0;
        for (int i = 0; i < N; i++) {
            result += coefficients[i] * values[i];
        }
        cout<<"result = "<<result<<'\n';
        return result == constant;
    }
};

#endif  // EQUALITY_H
