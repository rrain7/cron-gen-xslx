from get_data import gen_xlsx, write_date_2_excel_file


# 运行函数入口
if __name__ == '__main__':
    file_name = gen_xlsx()
    write_date_2_excel_file(file_name=file_name)
