#include "stdio.h"

// Tagged Unions are not supported in Go
// https://golang.org/doc/faq#variant_types

enum ShapeKind { Square, Rectangle, Circle};

struct Shape {
    int numOfSides;
    enum ShapeKind kind;  // using enum as tag for the union
    union {
        struct { int side; }; // square 
        struct { int length, breadth; }; // rectangle
        struct { int radius; }; // circle
    };
} phone;

// using getters and setters for accessing tagged unions guarantees safety

// check on the tag must happen before accessing the members
int getSquareSide(struct Shape* s) {
    if (s->kind == Square) {
        return s->side;
    }
    return 0;
}

// set on the tag should happen before assigning the members
void setSquareSide(struct Shape* s, int value) {
    s->kind = Square;
    s->side = value;
}

int getCircleRadius(struct Shape* s) {
    if (s->kind == Circle) {
        return s->radius;
    }
    return 0;
}

int main() {
    setSquareSide(&phone, 10);
    printf("side: %d length: %d breadth: %d \n", phone.side, phone.length, phone.breadth);

    // the tag phone.kind can be used to know, what fields of the union can be reliable

    return 0;
}