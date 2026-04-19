use Casbin\Enforcer;
use Casbin\Persist\Adapters\FileAdapter;

// way1
$e = new Enforcer('examples/basic_model.conf', 'examples/basic_policy.csv');

// way2
$a = new FileAdapter('examples/basic_policy.csv');
$e = new Enforcer('examples/basic_model.conf', $a);