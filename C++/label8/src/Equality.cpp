#include "Equality.h"
#include<bits/stdc++.h>
using namespace std;

template <typename Type>
Equality<Type>::Equality(int n,Type l)
{
    int kolvo=n;
    Type gran=l;
    Type elem;
    for (int i=0; i<n; i++){
        cin>>elem;
        a.push_back(elem-l);
    }
}

template <typename Type>
bool Equality<Type>::IsCorrect(Type h){
    for (auto el: a)
        if (el>gran || el<h)
            return false;
    return true;
}
