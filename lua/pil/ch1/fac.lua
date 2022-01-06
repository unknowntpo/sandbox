function fac (n)
    if n <= 1 then
        return n
    end
    return n * fac(n - 1)
end

for i = 0,10,1 do
 print(fac(i))
end
