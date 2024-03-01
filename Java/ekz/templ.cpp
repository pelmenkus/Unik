template
<typename T >
2 T t_max (T a , T b)
3
{
4 cout << " 1␣"; return a > b ? a : b;
5
}
6
7 template <>
8 int t_max <int >( int a , int b)
9
{
10 cout << " 2␣"; return a > b ? a : b;
11
}
12
13 int t_max (int a , int b) {
14 cout << " 3␣"; return a > b ? a : b;
15
}
16
17 int main ()
18
{
19 int x = 10 , y = 20;
20 cout << t_max (x ,y ); // Вывод : 3 20
21 return 0;
22
}
