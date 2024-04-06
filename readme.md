# The Binding of Isaac: Rebirth items & trinkets WS

Endpoints
---------

*   ### `/items`
    
    Returns all items.
    
    Parameters:
    
    *   `id` _(Optional)_
        
        Returns a single item matching the given item ID.
        
        _eg: `/items?id=3` returns the item "Spoon Bender"._
        
    *   `name` _(Optional)_
        
        Returns a single item matching the given name. Case insensitive.
        
        _eg: `/items?name=cRickEt's hEAd` returns the item "Cricket's Head"._
        
    *   `dlc` _(Optional)_
        
        Returns all items matching the given dlc. Case insensitive.
        
        _eg: `/items?dlc=repenTanCe` returns all repentance DLC items._
        
    *   `quality` _(Optional)_
        
        Returns all items matching the given quality.
        
        _eg: `/items?quality=4` returns all quality 4 items._
        
    *   `type` _(Optional)_
        
        Returns all items matching the given type. Type can be "Active" or "Passive". Case insensitive.
        
        _eg: `/items?type=acTive` returns all active items._
        
*   ### `/items/random`
    
    Returns a random item.
    
    Parameters:
    
    *   `qty` _(Optional)_
        
        Returns the number of random items given by the parameter. Items **do not** repeat.  
        If enough constraints are on the query so the number of results is lower than the given quantity, returns the max possible number of items matching the parameter constraints.
        
        _eg: `/items/random?qty=3` returns three random items._
        
    *   `dlc` _(Optional)_
        
        Returns a random item from the given DLC. Case insensitive.
        
        _eg: `/items/random?dlc=repenTanCe` returns a random item from the Repentance DLC. Case insensitive._
        
    *   `quality` _(Optional)_
        
        Returns a random item from the given quality.
        
        _eg: `/items/random?quality=4` returns a random quality 4 item._
        
    *   `type` _(Optional)_
        
        Returns a random item from the given type. Type can be "Active" or "Passive". Case insensitive. (Optional)
        
        _eg: `/items/random?type=acTive` returns a random Active item._
        

*   ### `/trinkets`
    
    Returns all trinkets.
    
    Parameters:
    
    *   `id` _(Optional)_
        
        Returns a single trinket matching the given trinket ID.
        
        _eg: `/trinkets?id=3` returns the trinket "AAA Battery"._
        
    *   `name` _(Optional)_
        
        Returns a single trinket matching the given name. Case insensitive.
        
        _eg: `/trinkets?name=swaLLoWed PenNy` returns the trinket "Swallowed Penny"._
        
    *   `dlc` _(Optional)_
        
        Returns all trinkets matching the given dlc. Case insensitive.
        
        _eg: `/trinkets?dlc=repenTanCe` returns all repentance DLC trinkets._
        
*   ### `/trinkets/random`
    
    Returns a random trinket.
    
    Parameters:
    
    *   `qty` _(Optional)_
        
        Returns the number of random trinkets given by the parameter. Trinkets **do not** repeat.  
        If enough constraints are on the query so the number of results is lower than the given quantity, returns the max possible number of trinkets matching the parameter constraints.
        
        _eg: `/trinkets/random?qty=3` returns three random trinkets._
        
    *   `dlc` _(Optional)_
        
        Returns a random item from the given DLC. Case insensitive.
        
        _eg: `/trinkets/random?dlc=repenTanCe` returns a random trinket from the Repentance DLC. Case insensitive._