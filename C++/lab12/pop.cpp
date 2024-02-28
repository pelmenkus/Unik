#include <bits/stdc++.h>

using namespace std;

class Iterator {
public:
    Iterator(int* row,int nado) : row_(row), n_(1), gcd_(0) {
        if (nado==1)
            update_gcd();
        else
            gcd_=0;
    }

    Iterator& operator++() {
        row_ += n_;
        update_gcd();
        return *this;
    }

    int operator*() const {
        return gcd_;
    }

    bool operator==(const Iterator& other) const {
        return row_ == other.row_;
    }

    bool operator!=(const Iterator& other) const {
        return !(*this == other);
    }

private:
    int* row_;
    int n_;
    int gcd_;
    void update_gcd() {
        //cout << abs(row_[0]) << " " << abs(row_[1]) << endl;
        gcd_ = __gcd(abs(row_[0]), abs(row_[1]));
        //cout << gcd_ << endl;

        //gcd=row_[n++];
    }
};

class vecIt{
private:
    vector <int> mas;
    int siz=0;
public:
    vecIt(){
    }
    add(int val){
        mas.push_back(val);
        siz++;
    }
    int* operator [](int ukaz){
        return &mas[ukaz];
    }
    Iterator begin(){
        return Iterator(&mas[0],1);
    }
    Iterator end(){
        return Iterator(&mas[siz-1],0);
    }
};

int main(){
    vecIt el;
    el.add(15);
    el.add(10);
    el.add(8);
    el.add(6);
    el.add(4);
    el.add(2);
    el.add(9);
    el.add(11);
    //cout<<el[3]<<'\n';
    vector <int> a;
    cout<<*el[3]<<" - mas[4]"<<'\n';
    for (auto ok=el.begin(); ok!=el.end(); ++ok){
        cout<<*ok<<'\n';
        int check=*ok;
        a.push_back(check);
    }
    vector<int> ans={5,2,2,2,2,1,1};
    if(a==ans)
        cout<<"TESTS Completed";
    else
        cout<<"ERROR";
    cout<<'\n';



    *el[0]=20;
    cout<<*el[0]<<" - mas[0] "<<'\n';
    a.clear();
    for (auto ok=el.begin(); ok!=el.end(); ++ok){
        cout<<*ok<<'\n';
        int check=*ok;
        a.push_back(check);
    }
    vector<int> ans2={10,2,2,2,2,1,1};
    if(a==ans2)
        cout<<"TESTS Completed";
    else
        cout<<"ERROR";


}
