#include <bits/stdc++.h>
using namespace std;

#include <cmath>
#include <vector>

template <typename T>
class Divs {
public:
    // Конструктор по умолчанию
    Divs() = default;

    // Конструктор с одним аргументом
    Divs(int num, int kek) : N(kek), divisors_(N, 0), primes(N) {

        decompose(num);
    }

    // Оператор присваивания с перегрузкой '*='
    Divs& operator*=(const Divs& other) {
        for (int i = 0; i < N; ++i) {
            divisors_[i] += other.divisors_[i];
            //cout<<divisors_[i]<<'\n';
        }
        return *this;
    }

    void Pusk(){
        for (auto el: divisors_){
            cout<<el<<"-ok"<<'\n';
        }
    }
    // Оператор умножения с перегрузкой '*'
    friend Divs operator*(const Divs& lhs, const Divs& rhs) {
        Divs result(lhs);
        result *= rhs;
        return result;
    }

    // Оператор перегрузки '&'
    friend Divs operator&(const Divs& lhs, const Divs& rhs) {
        Divs result(lhs);
        int params=1;
        for (int i = 0; i < lhs.N; ++i) {
            result.divisors_[i] = std::min(lhs.divisors_[i], rhs.divisors_[i]);
            //cout<<result.divisors_[i]<<'\n';
            if (result.divisors_[i]!=0)
                params*=pow(result.divisors_[i],i);
        }
        cout<<params<<'\n';
        return result;
    }

    // Операторы сравнения с перегрузкой '==', '!=', '<', '<=', '>', '>='
    friend bool operator==(const Divs& lhs, const Divs& rhs) {
        return lhs.divisors_ == rhs.divisors_;
    }
    friend bool operator!=(const Divs& lhs, const Divs& rhs) {
        return !(lhs == rhs);
    }
    friend bool operator<(const Divs& lhs, const Divs& rhs) {
        return lhs.divisors_ < rhs.divisors_;
    }
    friend bool operator<=(const Divs& lhs, const Divs& rhs) {
        return !(rhs < lhs);
    }
    friend bool operator>(const Divs& lhs, const Divs& rhs) {
        return rhs < lhs;
    }
    friend bool operator>=(const Divs& lhs, const Divs& rhs) {
        return !(lhs < rhs);
    }

    // Оператор преобразования к типу T
    explicit operator T() const {
        T result = 1;
        for (int i = 0; i < N; ++i) {
            result *= pow(primes[i], divisors_[i]);
        }
        return result;
    }

private:
    // Приватный метод, разложения числа на простые делители
    int N;
    std::vector<int> divisors_;
     std::vector<int> primes;

    void decompose(int num) {
        for (int i = 2; i * i <= num; ++i) {
            while (num % i == 0) {
                ++divisors_[i];
                num /= i;
            }
        }

        if (num != 1) {
            ++divisors_[num];
        }
    }
};

int main() {
    // Creating two Divs objects
    Divs<int> a(12, 5);
    Divs<int> b(20, 5);

    // Using the *= operator to multiply a and b
    b.Pusk();
    cout<<"divisors B"<<'\n';
    a.Pusk();
    cout<<"divisors A"<<'\n';
    a *= b;

    // Printing the result
    std::cout << "a *= b = " << static_cast<int>(a) << std::endl;
    cout<<"New divisors A"<<'\n';
    a.Pusk();
    // Using the & operator to find the common divisors of a and b
    Divs <int>c = a & b;

    // Printing the result
    std::cout << "Common divisors of a and b = " << static_cast<int>(c) << std::endl;

    return 0;
}
