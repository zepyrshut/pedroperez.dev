const sum = require('./index')

test('adds some numbers', () => {
  expect(sum(1, 2)).toBe(3)
  expect(sum(3, 9)).toBe(12)
})
