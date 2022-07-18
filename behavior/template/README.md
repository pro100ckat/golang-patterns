## Шаблонный метод
Это поведенческий паттерн, задающий скелет алгоритма в суперклассе и заставляющий подклассы реализовать конкретные шаги этого алгоритма.

В таком случае шаги конкретной операции одинаковы, но их реализация может отличаться. Такая ситуация подходит для использования паттерна Шаблонный метод.

Сперва мы определим базовый шаблонный алгоритм, который состоит из фиксированного количества методов. Это и будет нашим шаблонным методом. После этого мы реализуем методы для каждого шага, но шаблонный метод при этом трогать не будем.