package leetcode

/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val = 0, $next = null) {
 *         $this->val = $val;
 *         $this->next = $next;
 *     }
 * }
 */
//  class Solution {

//     /**
//      * @param ListNode $list1
//      * @param ListNode $list2
//      * @return ListNode
//      */
//     function mergeTwoLists($list1, $list2) {
//             $arr = [];
//             while($list1!=null){
//                 $arr[]=$list1->val;
//                 $list1 = $list1->next;
//                 while($list2!=null){
//                     $arr[] = $list2->val;
//                     $list2=$list2->next;
//                     break;
//                 }
//             }
//             while($list2){
//                 $arr[]=$list2->val;
//                 $list2=$list2->next;
//             }

//             sort($arr);
//             if($arr){
//                 $node = new ListNode($arr[0]);
//                 $bak = $node;
//             }else{
//                 return null;
//             }

//             foreach($arr as $k=>$val){
//                 if($k==0){
//                     continue;
//                 }
//                   $bak->next =  new ListNode($val);
//                   $bak = $bak->next;
//             }

//             return $node;
//     }
// }
