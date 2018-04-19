'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _classnames2 = require('classnames');

var _classnames3 = _interopRequireDefault(_classnames2);

var _isDefined = require('../utils/isDefined');

var _isDefined2 = _interopRequireDefault(_isDefined);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

var propTypes = {
  align: _propTypes2.default.oneOf(['top', 'middle', 'bottom']),
  className: _propTypes2.default.string,
  col: _propTypes2.default.number,
  order: _propTypes2.default.number,
  phone: _propTypes2.default.number,
  tablet: _propTypes2.default.number
};

var ROOT = 'mdc-layout-grid__cell';

var Cell = function Cell(_ref) {
  var _classnames;

  var align = _ref.align,
      className = _ref.className,
      col = _ref.col,
      order = _ref.order,
      phone = _ref.phone,
      tablet = _ref.tablet,
      otherProps = _objectWithoutProperties(_ref, ['align', 'className', 'col', 'order', 'phone', 'tablet']);

  return _react2.default.createElement('div', _extends({
    className: (0, _classnames3.default)(ROOT, (_classnames = {}, _defineProperty(_classnames, ROOT + '--span-' + col, (0, _isDefined2.default)(col)), _defineProperty(_classnames, ROOT + '--span-' + tablet + '-tablet', (0, _isDefined2.default)(tablet)), _defineProperty(_classnames, ROOT + '--span-' + phone + '-phone', (0, _isDefined2.default)(phone)), _defineProperty(_classnames, ROOT + '--order-' + order, (0, _isDefined2.default)(order)), _defineProperty(_classnames, ROOT + '--align-' + align, (0, _isDefined2.default)(align)), _classnames), className)
  }, otherProps));
};
Cell.propTypes = propTypes;
exports.default = Cell;