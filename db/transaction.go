package db

/*
命令模式
*/

// Command 执行数据库操作的命令接口
type Command interface {
	// Exec 执行insert、update、delete命令
	Exec() error
	// Undo 回滚命令
	Undo()
	// SetDb 设置关联的数据库
	setDb(db Db)
}

// 对于数据库来说，Transaction这个具体的结构体就是数据库用来控制事务的结构体 它控制的方式就是首先创建Command 命令切片，执行事务就是把传入的想要执行的命令存储Command 命令切片1 ，等到提交的时候，遍历存储好命令的切片1，在每一次的遍历中，调用Command本身去执行
// 然后把执行完的命令按照顺序存储到另外一个切片2， 在遍历切片1的每次循环中，一旦某次循环失败，那么就需要开始循环切片2，并调用Command 的Undo() 方法撤销操作
// Transaction Db事务实现，事务接口的调用顺序为begin -> exec -> exec > ... -> commit
type Transaction struct {
	name string
	db   Db
	cmds []Command
}

func NewTransaction(name string, db Db) *Transaction {
	return &Transaction{
		name: name,
		db:   db,
		cmds: nil,
	}
}

// Begin 开启一个事务
func (t *Transaction) Begin() {
	t.cmds = make([]Command, 0)
}

// Exec 在事务中执行命令，先缓存到cmds队列中，等commit时再执行
func (t *Transaction) Exec(cmd Command) error {
	if t.cmds == nil {
		return ErrTransactionNotBegin
	}
	cmd.setDb(t.db)
	t.cmds = append(t.cmds, cmd)
	return nil
}

// Commit 提交事务，执行队列中的命令，如果有命令失败，则回滚后返回错误
// 值得注意的是往往commit 本身就包含了如果执行命令失败就回滚的逻辑
func (t *Transaction) Commit() error {
	history := &cmdHistory{history: make([]Command, 0, len(t.cmds))}
	for _, cmd := range t.cmds {
		if err := cmd.Exec(); err != nil {
			history.rollback()
			return err
		}
		history.add(cmd)
	}
	return nil
}

/*
备忘录模式
   备忘录就主要是实现两个功能，
   一个是用自己的一个属性history-数组类型来记录存储一个个cmd结构体。
   第二个是当需要返回到起点的时候就是(从history数组的末尾开始倒序一个个调用cmd的方法undo 来撤销cmd已经执行过的事)
   一般都把后面括号里面的操作封装成一个rollback 函数来回滚
*/
// cmdHistory 命令执行历史
type cmdHistory struct {
	history []Command
}

func (c *cmdHistory) add(cmd Command) {
	c.history = append(c.history, cmd)
}

// 对已经添加的命令执行撤销操作 封装了一下
func (c *cmdHistory) rollback() {
	for i := len(c.history) - 1; i >= 0; i-- {
		c.history[i].Undo()
	}
}

// 因为插入和删除以及更新都会对数据库的内容进行修改所以只有这三个命令才会去实现命令的接口，实现事务
// InsertCmd 插入命令
// 这里是把一个单一的命令写一个结构体去控制这个单一的命令，并且给控制的结构体附加一个Undo() 方法
// Undo() 方法就是取消这个命令，如果不是用结构体控制这个插入的命令，我们就无法对这个命令进行撤销
// 如果想要给某个要执行的函数，附加一种让这个函数的执行行为回滚或者撤销那么你需要有个结构体去实现：
// 1. 我们原本要某个要执行的函数 2. 另外一个实现将刚刚执行的函数的结果删除的函数（这个在本质上就是撤销）
type InsertCmd struct {
	db         Db
	tableName  string
	primaryKey interface{}
	newRecord  interface{}
}

func NewInsertCmd(tableName string) *InsertCmd {
	return &InsertCmd{tableName: tableName}
}

func (i *InsertCmd) WithPrimaryKey(primaryKey interface{}) *InsertCmd {
	i.primaryKey = primaryKey
	return i
}

func (i *InsertCmd) WithRecord(record interface{}) *InsertCmd {
	i.newRecord = record
	return i
}

func (i *InsertCmd) Exec() error {
	return i.db.Insert(i.tableName, i.primaryKey, i.newRecord)
}

func (i *InsertCmd) Undo() {
	i.db.Delete(i.tableName, i.primaryKey)
}

func (i *InsertCmd) setDb(db Db) {
	i.db = db
}

// UpdateCmd 更新命令
type UpdateCmd struct {
	db         Db
	tableName  string
	primaryKey interface{}
	newRecord  interface{}
	oldRecord  interface{}
}

func NewUpdateCmd(tableName string) *UpdateCmd {
	return &UpdateCmd{tableName: tableName}
}

func (u *UpdateCmd) WithPrimaryKey(primaryKey interface{}) *UpdateCmd {
	u.primaryKey = primaryKey
	return u
}

func (u *UpdateCmd) WithRecord(record interface{}) *UpdateCmd {
	u.newRecord = record
	return u
}

func (u *UpdateCmd) Exec() error {
	if err := u.db.Query(u.tableName, u.primaryKey, u.oldRecord); err != nil {
		return err
	}
	return u.db.Update(u.tableName, u.primaryKey, u.newRecord)
}

func (u *UpdateCmd) Undo() {
	u.db.Update(u.tableName, u.primaryKey, u.oldRecord)
}

func (u *UpdateCmd) setDb(db Db) {
	u.db = db
}

// DeleteCmd 删除命令
type DeleteCmd struct {
	db         Db
	tableName  string
	primaryKey interface{}
	oldRecord  interface{}
}

func NewDeleteCmd(tableName string) *DeleteCmd {
	return &DeleteCmd{tableName: tableName}
}

func (d *DeleteCmd) WithPrimaryKey(primaryKey interface{}) *DeleteCmd {
	d.primaryKey = primaryKey
	return d
}

func (d *DeleteCmd) Exec() error {
	return d.db.Delete(d.tableName, d.primaryKey)
}

func (d *DeleteCmd) Undo() {
	d.db.Insert(d.tableName, d.primaryKey, d.oldRecord)
}

func (d *DeleteCmd) setDb(db Db) {
	d.db = db
}
