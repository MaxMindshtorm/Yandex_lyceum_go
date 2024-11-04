#include <iostream>
#include <string>

using namespace std;

struct Stack {
    char t;
    double n;
    Stack* next;
};

Stack* top1 = nullptr;
Stack* top2 = nullptr;

void push(double s, char a, Stack*& top) {
    Stack* p = (Stack*)malloc(sizeof(Stack));
    (*p).n = s;
    (*p).t = a;
    (*p).next = top;
    top = p;
}

Stack pop(Stack*& top) {
    Stack towp = *top;
    top = (*top).next;
    return towp;
}

int prior(char ch) {
    if (ch == '(') {
        return 0;
    }
    if (ch == '+' || ch == '-') {
        return 1;
    }
    if (ch == '*' || ch == '/') {
        return 2;
    }
    return -1;
}

void apply(Stack*& top, char z) {
    if (top == nullptr) {
        cout << "Error.";
        exit(0);
    }
    double b = (pop(top)).n;
    if (top == nullptr) {
        cout << "Error.";
        exit(0);
    }
    double a = (pop(top)).n;
    if (z == '+') {
        push(a + b, '@', top);
    } else if (z == '-') {
        push(a - b, '@', top);
    } else if (z == '*') {
        push(a * b, '@', top);
    } else {
        if (a == 0 || b == 0) {
            cout << "Error.";
            exit(0);
        }
        push(a / b, '@', top);
    }
}

int main() {
    double v = 0;
    string num = "";
    string s;
    cin >> s;
    int cnt = 0;
    for (auto i : s) {
        if (i == '(') {
            cnt++;
        }
        if (i == ')') {
            cnt--;
        }
    }
    if (cnt != 0) {
        cout << "Error.";
        exit(0);
    }
    for (int i = 0; i <= s.size(); i++) {
        if (s[i] >= '0' && s[i] <= '9') {
            while (s[i] >= '0' && s[i] <= '9') {
                num += s[i];
                i += 1;
            }
            v = stoi(num);
            num = "";
            push(v, '@', top1);
        }
        if (s[i] == '(') {
            push(0, s[i], top2);
        }
        if (s[i] == ')') {
            while ((*top2).t != '(') {
                apply(top1, (pop(top2)).t);
            }
            pop(top2);
        }
        if (s[i] == '-' && i == 0) {
            push(-1, '@', top1);
            push(0, '*', top2);
        } else if (s[i] == '-' && s[i - 1] == '(') {
            push(-1, '@', top1);
            push(0, '*', top2);
        } else if (s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/') {
            while (top2 != nullptr && prior(s[i]) <= prior((*top2).t)) {
                apply(top1, (pop(top2)).t);
            }
            push(0, s[i], top2);
        }
    }
    while (top2 != nullptr) {
        apply(top1, (pop(top2)).t);
    }
    cout << fixed;
    cout.precision(6);
    cout << (*top1).n;
    exit(0);
}
