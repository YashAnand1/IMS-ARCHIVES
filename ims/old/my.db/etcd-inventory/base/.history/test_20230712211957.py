import statistics
import matplotlib.pyplot as plt

GDP = [23.3, 17.7, 4.9, 4.3, 3.2, 3.1, 3, 2.1, 2, 1.8]
countries = ['USA', 'CHINA', 'JAPAN', 'GERMANY', 'INDIA', 'UK', 'FRANCE', 'ITALY', 'CANADA', 'SOUTH KOREA']

# calculating the median
median = statistics.median(GDP)

print("Median: ", median)

# plotting boxplot graph.
plt.bar(countries, GDP)

plt.ylabel("GDP")
plt.title("Top 10 Countries GDP on Bar Graph")
plt.xlabel("countries")
plt.xticks(rotation=45)
plt.axhline(y=GDP, color='r', linestyle='', label='Median GDP')
plt.legend()

plt.show()
