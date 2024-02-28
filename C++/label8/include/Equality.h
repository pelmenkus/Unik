#ifndef EQUALITY_H
#define EQUALITY_H
#include <bits/stdc++.h>
using namespace std;
template <typename Type>
class Equality
{
    int kolv;
    Type gran;
    vector<Type> a;
    public:
        Equality(int n, Type l);
       // virtual ~Equality();
        bool IsCorrect(Type h);
};

#endif // EQUALITY_H
