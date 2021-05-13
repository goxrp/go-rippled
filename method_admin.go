package gorippled

import (
	"github.com/grokify/simplego/database"
)

func StatusAndDebuggingMethods() []Method {
	return []Method{
		{
			Name:         MethodConsensusInfo,
			Category:     Category{Name: CategoryStatusAndDebugging},
			Summary:      "Get information about the state of consensus as it happens.",
			Description:  "The `consensus_info` command provides information about the [consensus process](https://xrpl.org/consensus.html) for debugging purposes.\n\n*The `consensus_info` method is an [admin method](https://xrpl.org/admin-rippled-methods.html) that cannot be run by unprivileged users.*",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodFeature,
			Category:     Category{Name: CategoryStatusAndDebugging},
			Summary:      "Get information about protocol amendments.",
			Description:  "The `feature` command returns information about [amendments](https://xrpl.org/amendments.html) this server knows about, including whether they are enabled and whether the server is voting in favor of those amendments in the [amendment process](https://xrpl.org/amendments.html#amendment-process).\n\nYou can use the `feature` command to configure the server to vote against or in favor of an amendment. This change persists even if you restart the server.\n\n*The `feature` method is an [admin method](https://xrpl.org/admin-rippled-methods.html) that cannot be run by unprivileged users.*",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodFetchInfo,
			Category:     Category{Name: CategoryStatusAndDebugging},
			Summary:      "Get information about the server's sync with the network.",
			Description:  "The `fetch_info` command returns information about objects that this server is currently fetching from the network, and how many peers have that information. It can also be used to reset current fetches.\n\n*The `fetch_info` method is an [admin method](https://xrpl.org/admin-rippled-methods.html) that cannot be run by unprivileged users.*",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodGetCounts,
			Category:     Category{Name: CategoryStatusAndDebugging},
			Summary:      "Get statistics about the server's internals and memory usage.",
			Description:  "The `get_counts` command provides various stats about the health of the server, mostly the number of objects of different types that it currently holds in memory.\n\n*The `get_counts` method is an [admin method](https://xrpl.org/admin-rippled-methods.html) that cannot be run by unprivileged users.*",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodPrint,
			Category:     Category{Name: CategoryStatusAndDebugging},
			Summary:      "Get information about internal subsystems.",
			Description:  "The `print` command returns the current status of various internal subsystems, including peers, the ledger cleaner, and the resource manager.\n\n*The `print` method is an [admin method](https://xrpl.org/admin-rippled-methods.html) that cannot be run by unprivileged users!*",
			IsOneOf:      true,
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodValidatorInfo,
			Category:     Category{Name: CategoryStatusAndDebugging},
			Summary:      "Get the server's validation settings, if configured as a validator.",
			Description:  "The `validator_info` method returns the current validator settings of the server, if it is configured as a validator.\n\n*The `validator_info` method is an [admin method](https://xrpl.org/admin-rippled-methods.html) that cannot be run by unprivileged users.*",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodValidatorListSites,
			Category:     Category{Name: CategoryStatusAndDebugging},
			Summary:      "Get information about sites that publish validator lists.",
			Description:  "The `validator_list_sites` command returns status information of sites serving validator lists.\n\n*The `validator_list_sites` method is an [admin method](https://xrpl.org/admin-rippled-methods.html) that cannot be run by unprivileged users!*",
			FunctionType: database.FunctionRead,
		},
		{
			Name:         MethodValidators,
			Category:     Category{Name: CategoryStatusAndDebugging},
			Summary:      "Get information about the current validators.",
			Description:  "The `validators` command returns human readable information about the current list of published and trusted validators used by the server.\n\n*The `validators` method is an [admin method](https://xrpl.org/admin-rippled-methods.html) that cannot be run by unprivileged users!*",
			FunctionType: database.FunctionRead,
		},
	}
}
