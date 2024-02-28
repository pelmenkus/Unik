#ifndef __CHISLO__H
#define __CHISLO__H

#include <bits/stdc++.h>
using namespace std;
class Chislo{
    int razmer,zpt;
    vector<int> a;
    public:
        Chislo(int d,int n,int p);
        ~Chislo();
        int countsize();
        int get_zn_index(int index);
        int* get_link_index(int index);
        int get_round(int index);
};
#endif // __CHISLO__H
