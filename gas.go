package lua

const (
	OP_MOVE_GAS         int = 1
	OP_MOVEN_GAS        int = 1
	OP_LOADK_GAS        int = 1
	OP_LOADBOOL_GAS     int = 1
	OP_LOADNIL_GAS      int = 1
	OP_GETUPVAL_GAS     int = 1
	
	OP_GETGLOBAL_GAS    int = 1
	OP_GETTABLE_GAS     int = 1
	OP_GETTABLEKS_GAS   int = 1
	
	OP_SETGLOBAL_GAS    int = 1
	OP_SETUPVAL_GAS     int = 1
	OP_SETTABLE_GAS     int = 1
	OP_SETTABLEKS_GAS   int = 1
	
	OP_NEWTABLE_GAS     int = 1
	
	OP_SELF_GAS         int = 1
	
	OP_ADD_GAS          int = 1
	OP_SUB_GAS          int = 1
	OP_MUL_GAS          int = 1
	OP_DIV_GAS          int = 1
	OP_MOD_GAS          int = 1
	OP_POW_GAS          int = 1
	OP_UNM_GAS          int = 1
	OP_NOT_GAS          int = 1
	OP_LEN_GAS          int = 1
	
	OP_CONCAT_GAS       int = 1
	
	OP_JMP_GAS          int = 1
	
	OP_EQ_GAS           int = 1
	OP_LT_GAS           int = 1
	OP_LE_GAS           int = 1
	
	OP_TEST_GAS         int = 1
	OP_TESTSET_GAS      int = 1
	
	OP_CALL_GAS         int = 1
	OP_TAILCALL_GAS     int = 1
	OP_RETURN_GAS       int = 1
	
	OP_FORLOOP_GAS      int = 1
	OP_FORPREP_GAS      int = 1
	
	OP_TFORLOOP_GAS     int = 1
	OP_SETLIST_GAS      int = 1
	
	OP_CLOSE_GAS        int = 1
	OP_CLOSURE_GAS      int = 1
	
	OP_VARARG_GAS       int = 1
	
	OP_NOP_GAS          int = 1
)