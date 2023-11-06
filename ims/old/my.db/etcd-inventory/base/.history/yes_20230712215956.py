import statistics
import matplotlib.pyplot as plt

gdp_value = [23.3, 17.7, 4.9, 4.3, 3.2, 3.1, 3, 2.1, 2, 1.8]
countries = ['USA', 'CHINA', 'JAPAN', 'GERMANY', 'INDIA',
             'UK', 'FRANCE', 'ITALY', 'CANADA', 'SOUTH KOREA']

# calculating the meadian
median = statistics.median(gdp_value)
print("Median:", median)

#ploting boxplot graph.
plt.bar(countries, gdp_value)
plt.title("Top 10 Countries - GDP on Bar Graph")
plt.xlabel("countries")                                                                                                                                                                                                                                                                                                                                                                                            
plt.ylabel("gdp trillion")
plt.xticks(rotation=45)
plt.axhline(y=gdp_value, color='r', linestyle='--', label='Median GDP')
plt.legend()
plt.show()