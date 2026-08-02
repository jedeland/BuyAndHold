import yfinance as yf

# Download data from 1927 to today
sp500 = yf.download('^KS11', start='1992-02-01')

# Save to a CSV file
sp500.to_csv('sources/kospi_historical.csv')