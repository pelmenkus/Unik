#include "chislo.h"
#include <bits/stdc++.h>
using namespace std;
Chislo::Chislo(int n, int d, int p){
    razmer=n-1;
    zpt=p;
    a.resize(n);
    for (int i=0; i<n; i++){
        a[i]=rand()%d;
    }
    a[p]=-228;
    for (int i=0; i<n; i++){
        cout<<a[i];
    }
    cout<<'\n';
}

Chislo::~Chislo(){
    a.clear();
}

int Chislo::countsize(){
    razmer=a.size()-1;
    return razmer;
}

int Chislo::get_zn_index(int index){
            if (index<razmer)
                return a[index];
            else
                return 0;
}

int* Chislo::get_link_index(int index){
            if (index<=razmer){
                int* link=&a[index];
                a[index]=0;
                return link;
            }
            else{
                if (index>zpt)
                    a.resize(razmer+index-zpt);
                else
                    a.resize(razmer+zpt-index);
                return &a[index];
            }
        }

int Chislo::get_round(int index){
    string s="";
    if (index<razmer)
        razmer=index;
    for (int i=0; i<razmer+1; i++){
        if (a[i]!=-228)
            s+=to_string(a[i]);
        else
            s+='.';
    }
    cout<<s<<'\n';
    float elem=stof(s);
    cout<<elem<<'\n';
    return 1;
}
