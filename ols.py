  
import pandas as pd
import statsmodels.api as sm
import os

def main():
    path:str = os.getcwd() + "/data.csv"

    # Matriz de datos globales
    data_df = pd.read_csv(path)
    data_df["indicador_madurez"] = 4 / data_df["madurez_succar"]
    data_df["const"] = 1

    # Matrices Y y X
    y_data = data_df["desv_costos"]
    x_data = data_df[["const", "indicador_madurez"]]

    regression = sm.OLS(endog=y_data, 
                        exog=x_data,
                        missing="drop")
    result_reg = regression.fit()
    result_summary = result_reg.summary()

    print(result_summary)
    print("\nDATA INGRESADA PARA LA ESTIMACIÓN\n")
    print(data_df)

    
if __name__ == "__main__":
    main()