assume cs: code, ds: data

data segment
string db 100, 99 dup (0)
dest db 100 dup("$")
data ends

code segment
println proc
    push bp ; запоминаем текущий bp
    mov bp, sp ; сохранение в bp вершины стека

    push ax
    push dx
    ; сохранили значения изменяемых в процедуре регситров

    mov dx, [bp + 4] ; получаем доступ к лежащему в стеке аргументу dx
    add dx, 2
    mov ah, 09h
    int 21h

    mov dl, 0Ah
    mov ah, 02h
    int 21h
    ; вывод символа строки

    pop ax
    pop dx
    ;восстановили значение регистров

    pop bp ; возвращаем старое значение в bp
    ret 2 ; выкидываем параметры процедуры из стека
println endp

strcpy proc
    push bp
    mov bp, sp; устанавливаем указатель стека в bp
    mov bx, [bp + 2] ; взяли адрес возрата
    mov si,[bp + 4]; string
    mov di,[bp + 6]; dest
    xor cx,cx
    mov cl, [si + 1] ;len(string)
    add cx, 2
    
    mov ax, di

    loop_start:
        mov dl, [si]
        mov [di], dl
        inc si
        inc di
        loop loop_start

    pop bp
    add sp, 6
    push ax
    push bx
    ret 
strcpy endp

start: 
    mov ax, data
    mov ds, ax

    mov dx, offset dest
    push dx
    xor dx, dx

    mov dx, offset string
    mov ah, 0Ah
    int 21h

    mov bl, [string + 1]; длина строки
    inc bl
    mov si, bx
    mov byte[string + si],'$'
    push dx ; кладём aдрес string в стек
    call strcpy
    ;на вершине стека находится указатель на начало строки dest можно написать pop dx, push dx
    call println
    mov ah, 4ch
    int 21h
code ends
end start