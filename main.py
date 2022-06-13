from get_data import gen_xlsx, write_date_2_excel_file, copy_and_rename


# 运行函数入口
if __name__ == '__main__':
    print("exec python script~\n")
    file_name = gen_xlsx()
    write_date_2_excel_file(file_name=file_name)
    copy_and_rename(file_name)
