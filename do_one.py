from get_data import gen_xlsx, write_date_2_excel_file, copy_and_rename, get_query_data_format


# 获取查询的数据时间
def get_query_time():
    y = int(input("请输入需要获取数据的年:\n"))
    m = int(input("请输入需要获取数据的月份:\n"))
    d = int(input("请输入需要获取数据的日期(天):\n"))
    h = int(input("请输入需要获取数据的起始小时 (起始时间， 范围 0- 23 ):\n"))
    return y, m, d, h


if __name__ == '__main__':
    print("此脚本仅能查询一次时间的数据, 若想获取多个时间段的数据， 请多次执行程序。。。")

    year, month, day, hour = get_query_time()
    query_time = get_query_data_format(year=year, month=month, day=day)
    # 生成excel文件
    file_name = gen_xlsx(month=month, day=day, hour=hour)
    # 将数据写入文件
    write_date_2_excel_file(file_name=file_name, year=year, month=month, day=day, hour=hour)
    copy_and_rename(file_name=file_name)

