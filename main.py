from get_data import gen_xlsx, write_date_2_excel_file, copy_and_rename, get_last_hour_time_info

# 运行函数入口
if __name__ == '__main__':
    print("exec python script~\n")
    year, month, day, hour = get_last_hour_time_info()
    file_name = gen_xlsx(month=month, day=day, hour=hour)
    write_date_2_excel_file(file_name=file_name, year=year, month=month, day=day, hour=hour)
    copy_and_rename(file_name)
