import statistics
import matplotlib.pyplot as plt

countries = ['USA', 'CHINA', 'JAPAN', 'GERMANY', 'INDIA', 'UK', 'FRANCE', 'ITALY', 'CANADA', 'SOUTH KOREA']
GDP = [23.3, 17.7, 4.9, 4.3, 3.2, 3.1, 3, 2.1, 2, 1.8]

# calculating the median
median = statistics.median(GDP)
print("Median:", median)

# plotting boxplot graph
plt.bar(countries, GDP)
plt.ylabel("GDP")
plt.title("Top 10 Countries GDP on Bar Graph")
plt.xlabel("countries")
plt.xticks(rotation=45)
plt.axhline(y=median, color='r', linestyle='', label='Median GDP')
plt.legend()

plt.show()

# import matplotlib.pyplot as plt   
# import statistics

# countries = ['USA', 'CHINA', 'JAPAN', 'GERMANY', 'INDIA',
#              'UK', 'FRANCE', 'ITALY', 'CANADA', 'SOUTH KOREA']
# gdp_values = [23.3, 17.7, 4.9, 4.3, 3.2, 3.1, 3, 2.1, 2, 1.8]

# # Calculate the mean
# median = statistics.median(gdp_values)
# print("Median:", median)

# # Plot the bar graph
# plt.bar(countries, gdp_values)
# plt.xlabel('Countries')
# plt.ylabel('GDP-Trillion USD')
# plt.title('GDP of Top 10 Countries (2023)')
# plt.xticks(rotation=45)
# plt.axhline(y=median, color='r', linestyle='--', label='Mean GDP')
# plt.legend()

# plt.show()