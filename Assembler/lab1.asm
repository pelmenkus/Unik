assume CS:code, DS:data

data segment
a db 10
b db 15
c db 4
d db 1
table db '0123456789ABCDEF'
int10 db "00000$"
int16 db "00000h$"
result db 0      ; Добавлено объявление переменной result

data ends

code segment
start:
    mov ax, data    ; загрузка сегмента данных
    mov ds, ax      ; установка ds в сегмент данных

    mov al, b
    mov bl, a
    add bl, d        ; a + d
    div bl           ; b / (a + d)

    ; Сохраняем результат деления в переменной result
    mov result, al

    mov al, c
    mov bl, 8
    mul bl
    add result, al   ; делить + умножить
    
    convert_to_int10:
        xor ax, ax
        mov al, result
        mov bl, 10
        mov si, offset int10 + 5 ; Здесь переводим указатель на $
        
        convert_to_int10_loop:
            div bl
            dec si
            mov di, offset table
            mov cx, di
            add cl, ah
            mov di, cx
            mov cl, [di]
            mov [si], cl
            xor ah, ah  ; Обнуляем остаток, чтоб делилось только частное
            test al, al ; Проверка частного на ноль
            jnz convert_to_int10_loop
        
        print_int10:
            mov ax, 0900h
            mov dx, si
            int 21h
            
            mov ax, 0200h
            mov dx, 10
            int 21h
    
    convert_to_int16:
        xor ax, ax
        mov al, result
        mov bl, 16
        mov si, offset int16 + 5 ; Здесь переводим указатель на $
        
        convert_to_int16_loop:
            div bl
            dec si
            mov di, offset table
            mov cx, di
            add cl, ah
            mov di, cx
            mov cl, [di]
            mov [si], cl
            xor ah, ah  ; Обнуляем остаток, чтоб делилось только частное
            test al, al ; Проверка частного на ноль
            jnz convert_to_int16_loop
        
        print_int16:
            mov ax, 0900h
            mov dx, si
            int 21h
    
    mov ax, 4C00h
    int 21h

code ends
end start
