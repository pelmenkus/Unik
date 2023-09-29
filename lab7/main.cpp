#include <bits/stdc++.h>
#include "chislo.h"

using namespace std;

int main(){
    Chislo elem(8,5,2);
    cout<<elem.countsize()<<'\n'<<elem.get_zn_index(3)<<'\n'<<elem.get_link_index(7)<<'\n';
    elem.get_round(9);
    return 0;
}
