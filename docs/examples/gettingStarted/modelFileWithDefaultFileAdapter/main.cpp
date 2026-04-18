#include <iostream>
#include <casbin/casbin.h>

int main() {
    // Create an Enforcer
    casbin::Enforcer e("path/to/model.conf", "path/to/policy.csv");

    // your code ..
}



// check permissions
    casbin::Enforcer e("../assets/model.conf", "../assets/policy.csv");

    if (e.Enforce({"alice", "/alice_data/hello", "GET"})) {
        std::cout << "Enforce OK" << std::endl;
    } else {
        std::cout << "Enforce NOT Good" << std::endl;
    }

    if (e.Enforce({"alice", "/alice_data/hello", "POST"})) {
        std::cout << "Enforce OK" << std::endl;
    } else {
        std::cout << "Enforce NOT Good" << std::endl;
    }