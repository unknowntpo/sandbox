# palindrome

num = 11
while true
  if num.to_s == num.to_s.reverse &&
      num.to_s(8) == num.to_s(8).reverse &&
      num.to_s(2) == num.to_s(2).reverse
    puts num
    break
  end
  # for search odd num only, we add 2 every time
  num += 2
end

